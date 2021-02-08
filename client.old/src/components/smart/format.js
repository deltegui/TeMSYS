export function formatSensorData(floatStr) {
  const float = parseFloat(floatStr);
  const oneDecimal = parseInt(float * 10);
  return oneDecimal / 10;
}