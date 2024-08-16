<template>
    <div class="home">
      <h1 class="text-2xl font-bold text-blue-600">Welcome to Your Vue.js App</h1>
      <p>{{ message }}</p>
    </div>
  </template>
  
  <script>
  import ApiService from '@/services/api';
  
  export default {
    name: 'HomeView',
    data() {
      return {
        message: ''
      }
    },
    mounted() {
      this.fetchData();
    },
    methods: {
      async fetchData() {
        try {
          const response = await ApiService.getExample();
          this.message = response.data.message;
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
    }
  }
  </script>