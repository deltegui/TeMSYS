<template>
  <hsc-window-style-metal style="position: fixed; z-index: 1">
    <hsc-window
      key="data-selector"
      :title="winTitle"
      :resizable="true"
      :closeButton="true"
      :isOpen.sync="isOpen"
      :width="1000"
      :height="600"
    >
      <canvas :id="`show-chart-${name}`"></canvas>
    </hsc-window>
  </hsc-window-style-metal>
</template>

<script>
import api from "../../../api";
import {ReportsChart} from '../../../charts';

export default {
  name: "data-viewer",
  props: {
    close: Function,
    showdata: Object
  },
  data: function() {
    return {
      isOpen: true,
      name: 0,
      winTitle: '',
    };
  },
  watch: {
    isOpen: function() {
      this.close();
    }
  },
  methods: {
    generateRandomName() {
      const arr = new Uint32Array(10);
      window.crypto.getRandomValues(arr);
      const str = arr.reduce((prev, current) => `${prev}${current}`, "");
      return str + this.showdata.sensor;
    },
    setWindowTitle() {
      const {
        sensor,
        date,
      } = this.showdata;
      const dateFromStr = formatDate(date.from);
      const dateToStr = formatDate(date.to);
      this.winTitle = `${sensor} from ${dateFromStr} to ${dateToStr}`;
    },
  },
  mounted() {
    this.setWindowTitle();
    this.name = this.generateRandomName();
    const chart = new ReportsChart(`show-chart-${this.name}`);
    api.report
      .getByDateRange({
        name: this.showdata.sensor,
        fromDate: this.showdata.date.from,
        toDate: this.showdata.date.to,
      })
      .then(reports => reports.filter(rep => this.showdata.reports.includes(rep.type)))
      .then(reports => {
        chart.data = reports;
        chart.draw();
      });
  }
};

function formatDate(date) {
  return `${date.getDate()}/${date.getMonth()+1}/${date.getFullYear()}`;
}
</script>