<script setup lang="ts">
import { ref } from 'vue';

const searchTerm = ref<string>("");
const emit = defineEmits(['onSearch']);

const handleSearch = () => {
  emit('onSearch', searchTerm.value.trim());
};
</script>

<template>
  <div class="px-4 w-full max-w-3xl mx-auto">
    <div
      class="flex items-center w-full h-12 rounded-lg bg-white shadow-md hover:shadow-lg transition-shadow duration-200">
      <!-- Search Input -->
      <input v-model="searchTerm" type="search" name="search" aria-label="Search" placeholder="Search..."
        @keyup.enter="handleSearch"
        class="flex-1 bg-transparent px-4 text-gray-700 placeholder-gray-400 focus:outline-none" />

      <!-- Search Button -->
      <button type="submit" @click="handleSearch" aria-label="Submit Search"
        class="flex items-center justify-center px-4 h-full text-white bg-blue-600 hover:bg-blue-700 rounded-r-lg transition-colors">
        <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
input[type="search"]::-webkit-search-decoration,
input[type="search"]::-webkit-search-cancel-button,
input[type="search"]::-webkit-search-results-button,
input[type="search"]::-webkit-search-results-decoration {
  -webkit-appearance: none;
}
</style>
