<script setup lang="ts">
import type { Email } from '../models/Email';
import type { EmailAddress } from '@/models/Address';
import type { Hit } from '../models/Hit';
import { ref } from 'vue';

const email = ref<Email | null>(null);

const displayEmail = (h: Hit) => {
    email.value = h._source.email;
}

const showTo = (to: EmailAddress[]): string => {
    if (to === undefined || to === null || to?.length === 0) {
        return 'N/A';
    } else if (to?.length === 1) {
        return to[0].Address;
    }
    return to.map(email => email?.Address).join(', ');
}

defineExpose({
    displayEmail
});

</script>

<template>
    <div class="overflow-y-auto overflow-hidden w-full h-full">
        <div class="p-6 overflow-scroll px-0 h-full">
            <div v-if="email !== null" class="py-10 px-8 space-y-4 ">
                <p>
                    <strong>Subject:</strong> {{ email["Subject"] }}
                </p>
                <p>
                    <strong>From:</strong> {{ email["From"] }}
                </p>
                <p>
                    <strong>To:</strong> {{ showTo(email["To"]) }}
                </p>
                <pre>{{ email["Body"] }}</pre>
            </div>
            <div v-else class="h-screen flex">
                <div class="m-auto">
                    <p>The email content will be displayed here.</p>
                </div>
            </div>
        </div>
    </div>
</template>
