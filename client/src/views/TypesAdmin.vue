<template>
  <div class="container">
    <h1>ReportTypes Admin</h1>
    <h2>Create</h2>
    <input class="temsys-input" v-model="newType" type="text" placeholder="Report Type"/>
    <button class="temsys-btn temsys-green" @click="createReportType">Create</button>
    <div v-if="error" class="temsys-input">
      {{error}}
    </div>
    <h2>Types list</h2>
    <p v-for="type in reportTypes" v-bind:key="type">{{type}}</p>
    <h4 v-if="reportTypes.length === 0">There is no ReportTypes!</h4>
  </div>
</template>

<script lang="ts">
import { reportService } from '@/services';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'TypeAdmin',
  data(): {
  reportTypes: string[];
  error: string;
  newType: string;
  } {
    return {
      reportTypes: [],
      error: '',
      newType: '',
    };
  },
  mounted() {
    this.updateList();
  },
  methods: {
    createReportType() {
      reportService.saveReportType(this.newType)
        .then(this.updateList.bind(this))
        .then(() => {
          this.newType = '';
        })
        .catch(() => {
          this.error = `ReportType: ${this.newType} already exists`;
        });
    },

    updateList() {
      reportService.getAllReportTypes()
        .then((types) => {
          this.reportTypes = types;
        });
    },
  },
});
</script>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  padding: 20px 20vw 20px 20vw;
  margin-top: 40px;

  display: flex;
  flex-direction: column;
}

@media only screen and (max-width: 900px) {
  .container {
    padding: 20px;
  }
}

.container > h1 {
  width: 100%;
}
</style>
