export class DataAsset {

  constructor(data) {
    this.data = data;
  }

  max() {
    return this.data.reduce((prev, current) => {
      return current >= prev ? current : prev;
    }, Number.MIN_VALUE);
  }

  min() {
    return this.data.reduce((prev, current) => {
      return current <= prev ? current : prev;
    }, Number.MAX_VALUE);
  }

  average() {
    return calculateAverageFor(this.data);
  }

  standardDeviation() {
    return Math.sqrt(this.variance());
  }

  variance() {
    const average = this.average();
    const distanceToAverage = this.data.map(v => Math.pow(v - average, 2));
    return calculateAverageFor(distanceToAverage);
  }

  coefficientVariation() {
    return (this.standardDeviation() / Math.abs(this.average()));
  }

}

export function calculateAverageFor(data) {
  const sum = data.reduce((sum, val) => sum + val, 0);
  const len = data.length;
  return sum / len;
}