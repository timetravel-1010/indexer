<script setup lang="ts">
import EmailInbox from './components/EmailInbox.vue';
import MainHeader from './components/MainHeader.vue';

import { ref } from 'vue';

const emailInbox = ref();

const setTerm = (term: string) => {
  emailInbox.value.searchEmails(term);
};
</script>

<template>
  <div class="h-screen flex flex-col">
    <!-- Header: Responsive Spacing and Layout -->
    <header>
      <MainHeader @on-search="setTerm" class="px-4 sm:px-6 md:px-8" />
    </header>

    <!-- Main Content Area: Adjust Layout for Small and Large Screens -->
    <div class="flex-1 flex flex-col lg:flex-row overflow-x-hidden">
      <!-- Email Inbox -->
      <main class="flex-1 w-full p-4 sm:p-6 lg:p-8 space-y-6">
        <EmailInbox ref="emailInbox" />
      </main>
    </div>
  </div>
</template>

<style></style>
