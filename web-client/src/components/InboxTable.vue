<script setup lang="ts">

import type { Hit } from '@/models/Hit';
import type { EmailAddress } from '@/models/Address';

const props = defineProps<{
    emails: Hit[];
}>()



const showTo = (to: EmailAddress[]): string => {

    if (to === undefined || to === null || to?.length === 0) {
        return 'N/A';
    } else if (to?.length === 1) {
        return to[0].Address;
    }
    return to[0].Address + ', ...';
}


</script>

<template>
    <div class="overflow-y-auto overflow-hidden w-full h-full">
        <div class="px-4 py-2 flex items-center justify-between border-l border-r border-b">
            Results
        </div>
        <div class="pt-3 pb-4 h-4/5">
            <div class="p-6 overflow-scroll px-0 h-full">
                <table class="mt-4 w-full min-w-max table-auto text-left">
                    <thead>
                        <tr>
                            <th
                                class="border-y border-blue-gray-100 bg-blue-gray-50/50 p-4 transition-colors hover:bg-blue-gray-50">
                                <p
                                    class="antialiased font-sans text-sm text-blue-gray-900 flex items-center justify-between gap-2 font-normal leading-none opacity-70">
                                    From
                                </p>
                            </th>
                            <th
                                class="border-y border-blue-gray-100 bg-blue-gray-50/50 p-4 transition-colors hover:bg-blue-gray-50">
                                <p
                                    class="antialiased font-sans text-sm text-blue-gray-900 flex items-center justify-between gap-2 font-normal leading-none opacity-70">
                                    To </p>
                            </th>
                            <th
                                class="border-y border-blue-gray-100 bg-blue-gray-50/50 p-4 transition-colors hover:bg-blue-gray-50">
                                <p
                                    class="antialiased font-sans text-sm text-blue-gray-900 flex items-center justify-between gap-2 font-normal leading-none opacity-70">
                                    Subject
                                </p>
                            </th>
                            <th
                                class="border-y border-blue-gray-100 bg-blue-gray-50/50 p-4 transition-colors hover:bg-blue-gray-50">
                                <p
                                    class="antialiased font-sans text-sm text-blue-gray-900 flex items-center justify-between gap-2 font-normal leading-none opacity-70">
                                    Date
                                </p>
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr @click="$emit('onSelect', email)" v-for="email in props.emails" :key="email._id"
                            class="cursor-pointer hover:bg-gray-200 ">
                            <td class="p-4 border-b border-blue-gray-50">
                                <div class="flex items-center gap-3">
                                    <div class="flex flex-col">
                                        <p
                                            class="block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-normal">
                                            {{ email._source.email["From"] }}
                                        </p>
                                    </div>
                                </div>
                            </td>
                            <td class="p-4 border-b border-blue-gray-50">
                                <div class="flex items-center gap-3">
                                    <div class="flex flex-col">
                                        <p
                                            class="block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-normal">
                                            {{ showTo(email._source.email["To"]) }}
                                        </p>
                                    </div>
                                </div>
                            </td>
                            <td class="p-4 border-b border-blue-gray-50">
                                <div class="flex flex-col">
                                    <p
                                        class="block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-normal">
                                        {{ email._source.email["Subject"] }}
                                    </p>
                                </div>
                            </td>
                            <td class="p-4 border-b border-blue-gray-50">
                                <p
                                    class="block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-normal">
                                    {{ email._source.email["Date"] }}
                                </p>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
