<script setup lang="ts">
import type { Email } from '../models/Email';
import type { EmailAddress } from '@/models/Address';
import type { Hit } from '../models/Hit';
import { ref } from 'vue';

const email = ref<Email | null>(null);

const displayEmail = (h: Hit) => {
  email.value = h._source.email;
};

const showTo = (to: EmailAddress[]): string => {
  if (to === undefined || to === null || to?.length === 0) {
    return 'N/A';
  } else if (to?.length === 1) {
    return to[0].Address;
  }
  return to.map((email) => email?.Address).join(', ');
};

defineExpose({
  displayEmail,
});
</script>

<template>
  <div class="overflow-y-auto w-full h-full">
    <div class="p-6 sm:p-8 lg:px-16 overflow-auto h-full">
      <div v-if="email !== null"
        class="py-10 px-6 sm:px-8 lg:px-12 space-y-6 bg-gray-50 dark:bg-gray-800 rounded-lg shadow-md">
        <p class="text-lg sm:text-xl font-semibold text-gray-700 dark:text-gray-300">
          <strong>Subject:</strong> {{ email["Subject"] }}
        </p>
        <p class="text-gray-600 dark:text-gray-400">
          <strong>From:</strong> {{ email["From"] }}
        </p>
        <p class="text-gray-600 dark:text-gray-400">
          <strong>To:</strong> {{ showTo(email["To"]) }}
        </p>
        <pre
          class="bg-gray-100 dark:bg-gray-900 p-4 rounded text-sm sm:text-base text-gray-800 dark:text-gray-100 overflow-auto">
          {{ email["Body"] }}
        </pre>
      </div>
      <div v-else class="flex items-center justify-center h-full text-center">
        <p class="text-gray-500 dark:text-gray-400 text-lg">
          The email content will be displayed here.
        </p>
      </div>
    </div>
  </div>
</template>
