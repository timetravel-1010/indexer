<script setup lang="ts">
import { ref, computed } from 'vue';
import type { Email } from '../models/Email';
import InboxTable from './InboxTable.vue';
import ContentPane from './ContentPane.vue';
import type { EmailResponse } from '../models/EmailResponse';
import type { Hit } from '../models/Hit';

const emails = ref<Hit[]>([]);
const isLoading = ref(false); // Loading state
const contentPane = ref();
const error = ref<string | null>(null); // Error message

const onDisplayEmail = (e: Email) => {
  contentPane.value.displayEmail(e);
};

const searchEmails = (term: string) => {
  isLoading.value = true; // Start loading
  error.value = null; // Reset error
  let url = new URL('http://localhost:8080/emails/search?index=custom&page=10');
  url.searchParams.append('term', term);
  fetch(url, {
    method: 'GET',
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then((data: EmailResponse) => {
      emails.value = data.hits.hits;
      console.log('data: ', data.hits);
    })
    .catch((error) => {
      console.error('Error: ', error);
      error.value = 'Failed to load emails. Please try again.';
    })
    .finally(() => {
      isLoading.value = false; // Stop loading
    });
};

const hasEmails = computed(() => emails.value.length > 0);

defineExpose({
  searchEmails,
});
</script>

<template>
  <div class="flex flex-col lg:flex-row w-full h-full">
    <!-- Left Pane: Inbox Table -->
    <div class="flex flex-col w-full lg:w-1/2 p-4 space-y-4 overflow-y-auto">
      <div v-if="error" class="text-red-500">
        {{ error }}
      </div>
      <InboxTable v-if="hasEmails" @on-select="onDisplayEmail" :emails="emails" />
      <p v-else class="text-gray-500">No emails found.</p>
      <button class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded" @click="searchEmails('')"
        :disabled="isLoading">
        {{ isLoading ? "Loading..." : "Load More" }}
      </button>
    </div>

    <!-- Right Pane: Content Pane -->
    <div class="flex-1 p-4 overflow-y-auto">
      <ContentPane ref="contentPane" />
    </div>
  </div>
</template>
