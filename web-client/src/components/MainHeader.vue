<script setup lang="ts">
import SearchBar from './SearchBar.vue';
import type { Email } from '../models/Email';

const props = defineProps<{
    emails?: Email[]
}>()

defineEmits(['onSearch']);

</script>


<template>
    <div class="main-header flex h-20 font-poppins font-bold items-center p-5">
        <img src="../assets/mail-box.png" alt="mailbox-icon" class="ml-10" width="4%">
        <h1 class="text-4xl ml-4">
            ZincMail
        </h1>
        <SearchBar :emails=props.emails />
    </div>
</template>


<style>
.main-header {
    background-color: #11009E;
    color: white;
}

.center-vertically {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
}
</style>
