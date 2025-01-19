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
}

const searchEmails = (term: string) => {
    //console.log("term:", term);
    let url = new URL('http://localhost:8080/emails/search?index=custom&page=10');
    url.searchParams.append("term", term);
    fetch(url, {
        method: 'GET',
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then((data: EmailResponse) => {
            emails.value = data.hits.hits;
            console.log('data: ', data.hits);
        })
        .catch(error => {
            console.error('Error: ', error);
        });
};

defineExpose({
    searchEmails
});

</script>


<template>
    <div class="flex flex-col w-1/2">
        <InboxTable @on-select="onDisplayEmail" :emails=emails />
        <button>Load more</button>
    </div>
    <ContentPane ref="contentPane" />
</template>


