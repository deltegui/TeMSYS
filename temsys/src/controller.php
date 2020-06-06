<?php

namespace Temsys\Controllers;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Twig\Environment;
use Slim\Csrf\Guard;

class Controller {

    private Environment $twig;

    private Response $response;

    private Request $request;

    private Session $session;

    private static ?Guard $csrf = null;

    public function __construct() {
        $loader = new \Twig\Loader\FilesystemLoader('../templates');
        $this->twig = new \Twig\Environment($loader);
    }

    public static function setCsrf(Guard $csrf) {
        if(self::$csrf == null) {
            self::$csrf = $csrf;
        }
    }

    protected function render(string $view, array $data = []) {
        $this
            ->response
            ->getBody()
            ->write($this->twig->render($view, $data));
    }

    public function session(): Session {
        return $this->session;
    }

    public function csrfData(): array {
        $csrf = self::$csrf;
        $nameKey = $csrf->getTokenNameKey();
        $valueKey = $csrf->getTokenValueKey();
        $name = $this->request->getAttribute($nameKey);
        $value = $this->request->getAttribute($valueKey);
        return [
            'nameKey' => $nameKey,
            'name' => $name,
            'value' => $value,
            'valueKey' => $valueKey,
        ];
    }

    /**
     * A way to hide the three params of a request in Slim into
     * this class. Simplifies controller methods and let me
     * easilly create a small api for them.
     * To use this, when you map a controller method you must
     * add an underscore in front of it. So if your method
     * is called 'hello', you must map it as '_hello'
     */
    public function __call(string $method, array $args) {
        [$request, $response, $callArgs] = $args;
        $this->request = $request;
        $this->response = $response;
        $realMethod = ltrim($method, '_');
        call_user_func([$this, $realMethod], $callArgs);
        return $this->response;
    }
}

class Session {

    public function set($key, $value) {
        $_SESSION[$key] = $value;
    }

    public function get($key) {
        return $_SESSION[$key];
    }

    public function flush() {
        session_unset();
    }

    public function destroy() {
        session_destroy();
    }
}
