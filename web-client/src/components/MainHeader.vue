<script setup lang="ts">
import SearchBar from './SearchBar.vue';
import type { Email } from '../models/Email';

const props = defineProps<{
  emails?: Email[]
}>();

defineEmits(['onSearch']);
</script>

<template>
  <header class="bg-[#11009E] text-white">
    <div class="max-w-7xl mx-auto">
      <div class="flex flex-col sm:flex-row items-center gap-4 sm:gap-8 p-4 sm:p-6">
        <!-- Logo Section -->
        <div class="flex items-center gap-3 sm:gap-4">
          <img src="../assets/mail-box.png" alt="ZincMail logo" class="w-8 h-8 sm:w-10 sm:h-10 object-contain" />
          <h1 class="text-2xl sm:text-3xl lg:text-4xl font-semibold">
            ZincMail
          </h1>
        </div>

        <!-- Search Section -->
        <div class="w-full sm:flex-1 max-w-2xl">
          <SearchBar :emails="emails" />
        </div>
      </div>
    </div>
  </header>
</template>
