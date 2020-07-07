// Use this defines to configure your hardware
//#define ETHERNET //Arduino's ethernet shield
//#define DHT11_SENSOR // DHT11 Termometer and Humidity sensor.
//#define TERMOMETER // Only termometer sensor
//#define NEOMCU // NeoMCU V3 board

#define ERROR_PIN 2
#define INFO_PIN 0

#ifdef ETHERNET
#include <SPI.h>
#include <Ethernet.h>
#endif

#ifdef DHT11_SENSOR
#include <DHT.h>
#endif

#ifdef NEOMCU
#include "Arduino.h"
#include <ESP8266WiFi.h>
#define WIFINAME "hola"
#define WIFIPASS "hace calor"
#endif

#ifdef TERMOMETER
  #define VOLTS 3.3
#endif

#define SERIALFREQ 9600

class JSONSerialize {
  public:
    virtual String toJSON() {}
};

class Sender {
  public:
    virtual void sendJSON(JSONSerialize *j) {}
};

class Termometer {
  public:
    virtual float get_temperature() {}
};

class Humidity {
  public:
    virtual float get_humidity() {}
};

#ifdef DHT11_SENSOR
class DHTSensor: public Humidity, public Termometer, public JSONSerialize {
  private:
    DHT *d;
  public:
    DHTSensor(int digitalPin) {
      this->d = new DHT(digitalPin, DHT11);
      this->d->begin();
    }

    ~DHTSensor() {
      delete this->d;
    }

    float get_temperature() {
      return this->d->readTemperature() + 4.0;
    }

    float get_humidity() {
      return this->d->readHumidity();
    }

    String toJSON() {
      return String("{\"temperature\": ") + this->get_temperature() + String(", \"humidity\": ") + this->get_humidity() + String("}");
    }
};
#endif

#ifdef TERMOMETER
class TempSensor : public Termometer, public JSONSerialize {
  private:
    int pin;
  public:
    TempSensor(int analogPin) {
      this->pin = analogPin;
    }

    float get_temperature() {
      int reading = analogRead(this->pin);
      float voltage = reading * VOLTS * 100;
      voltage /= 1024.0;
      return voltage;
    }

    String toJSON() {
      return String("{\"temperature\": ") + this->get_temperature() + String("}");
    }
};
#endif

#ifdef ETHERNET
class EthernetSender: public Sender {
  private:
    EthernetServer server;
    IPAddress ip;
    byte mac[6];

  public:
    EthernetSender(): EthernetSender(192, 168, 1, 177) {}

    EthernetSender(int one, int two, int third, int four): server(80), ip(one, two, third, four), mac{0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED} {
      Ethernet.begin(this->mac, this->ip);
      this->server.begin();
    }

    void sendJSON(JSONSerialize *t) {
      EthernetClient client = this->server.available();
      if (client) {
        String json = t->toJSON();
        Serial.println("new client");
        client.println("HTTP/1.1 200 OK");
        client.println("Content-Type: application/json;charset=utf-8");
        client.println("Server: ArduingoEthernet");
        client.println("Access-Control-Allow-Origin: *");
        client.println(String("Content-Length: ") + json.length());
        client.println("");
        client.println(json);
        client.stop();
      }
    }
};
#endif

#ifdef NEOMCU
class NeomcuSender: public Sender {
  private:
    WiFiServer server;

    void printMacAdress() {
      byte mac[6]; 
      WiFi.macAddress(mac);
      Serial.print("MAC: ");
      for(int i = 5; i >= 0; i--) {
        Serial.print(mac[i],HEX);
        if(i != 0) {
          Serial.print(":");
        }
      }
    }
  public:
    NeomcuSender(): server(80) {
      WiFi.begin(WIFINAME, WIFIPASS);
      while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        digitalWrite(ERROR_PIN, LOW);
        Serial.println("NOT CONNECTED. RETRIYING...");
      }
      Serial.println("CONNECTED!");
      Serial.println(WiFi.localIP());
      this->printMacAdress();
      digitalWrite(ERROR_PIN, HIGH);
      this->server.begin();
    }

    void sendJSON(JSONSerialize *t) {
      WiFiClient client = this->server.available(); 
      if(client) {
        Serial.println("We have a client!!!");
        String json = t->toJSON();
        client.println("HTTP/1.1 200 OK");
        client.println("Content-Type: application/json");
        client.println("Server: ArduingoWiFi");
        client.println("Access-Control-Allow-Origin: *");
        client.println(String("Content-Length: ") + json.length());
        client.println("");
        client.println(json);
        delay(1);
        client.stop();
      }
    }
};
#endif

class ErrorNotifier {
  private:
    int pin;

    void notifyPattern(int pattern[], int length) {
      for (int i = 0; i < length; i++) {
        delay(500);
        digitalWrite(this->pin, pattern[i]);
      }
      delay(4000);
    }

  public:
    ErrorNotifier(int pin): pin(pin) {
      pinMode(this->pin, OUTPUT);
      Serial.begin(SERIALFREQ);
    }

    void nullError() {
      int pattern[] = {1, 0, 1, 0};
      this->notifyPattern(pattern, 4);
      Serial.println("Parece que algo es NULL. La cosa no funciona");
    }
};

JSONSerialize *serializer;
Sender *sender;
ErrorNotifier notifier(ERROR_PIN);

void setup() {
#ifdef DHT11
  serializer = new DHTSensor(INFO_PIN);
#endif
#ifdef TERMOMETER
  serializer = new TempSensor(INFO_PIN);
#endif
#ifdef ETHERNET
  sender = new EthernetSender();
#endif
#ifdef NEOMCU
  sender = new NeomcuSender();
#endif
}

void loop() {
  if (sender == NULL || serializer == NULL) {
    notifier.nullError();
  }
  sender->sendJSON(serializer);
}
