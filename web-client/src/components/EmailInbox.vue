<script setup lang="ts">
import { ref } from 'vue';
import type { Email } from '../models/Email';
import InboxTable from './InboxTable.vue';
import ContentPane from './ContentPane.vue';
import type { EmailResponse } from '../models/EmailResponse';
import type { Hit } from '../models/Hit';

const emails = ref<Hit[]>([]);
const contentPane = ref();

const onDisplayEmail = (e: Email) => {
  contentPane.value.displayEmail(e);
};

const searchEmails = async (term: string) => {
  let url = new URL('http://localhost:8080/emails/search?index=custom&page=10');
  url.searchParams.append('term', term);
  const response = await fetch(url, {
    method: 'GET',
  });

  if (!response.ok) {
    console.error('Error in request: ', response);
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const data = await response.json() as EmailResponse;
  emails.value = data.hits.hits;
  console.log('data: ', data.hits);
};

defineExpose({ searchEmails });
</script>

<template>
  <div class="flex flex-col lg:flex-row w-full h-full">
    <!-- Left Pane: Inbox Table -->
    <div class="flex flex-col w-full lg:w-1/2 p-4 space-y-4">
      <InboxTable @on-select="onDisplayEmail" :emails="emails" />
      <button class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded">
        Load More
      </button>
    </div>

    <!-- Right Pane: Content Pane -->
    <div class="flex-1 p-4">
      <ContentPane ref="contentPane" />
    </div>
  </div>
</template>
