<script setup lang="ts">
import EmailInbox from './components/EmailInbox.vue';
import MainHeader from './components/MainHeader.vue';

import { ref } from 'vue';

const emailInbox = ref();

const setTerm = (term: string) => {
    emailInbox.value.searchEmails(term);
}


</script>

<template>
    <div class="h-screen flex flex-col">
        <header>
            <MainHeader @on-search="setTerm" />
        </header>
        <div class="flex-1 flex overflow-x-hidden">
            <main class="flex w-full">
                <EmailInbox ref="emailInbox" />
            </main>
        </div>
    </div>
</template>

<style></style>
