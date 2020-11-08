<template>
    <v-app>
        <TopBar/>
        <v-content>
            <v-layout fill-height>
                <router-view/>
            </v-layout>
        </v-content>
    </v-app>
</template>

<script lang="ts">
    import Vue from 'vue';
    import TopBar from '@/components/TopBar.vue';

    export default Vue.extend({
        name: 'App',

        components: {
            TopBar,
        },
        data: () => ({}),
    });
</script>
