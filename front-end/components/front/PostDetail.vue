<template>
<div>
    <v-container>
        <v-row>
            <v-col xl="12" lg="12" md="12" sm="12">
                <div class="my-10">
                    <div class="my-2">
                        <h2>{{ title }}</h2>
                    </div>
                    <v-card-text v-for="(item, index) in tags" :key="index">
                        <nuxt-link :to="{ path: '/blog?search=' + item.name }">
                            <v-chip>{{ item.name }}</v-chip>
                        </nuxt-link>
                    </v-card-text>
                    <div class="mb-3">
                        <div v-html="content" id="app">
                        </div>
                    </div>
                </div>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
export default {
    name: "PostDetailComp",
    props: {
        title: String,
        content: String,
        tags: Array,
    },
};
</script>

<style >
</style>
