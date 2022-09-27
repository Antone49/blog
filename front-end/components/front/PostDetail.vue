<template>
<div>
    <v-container>
        <v-row>
            <v-col xl="8" lg="8" md="8" sm="12">
                <div class="my-10">
                    <img style="width: 100%" :src="'/images/' + image" :alt="title" />
                    <div class="my-2">
                        <h2>{{ title }}</h2>
                    </div>
                    <div class="mb-3">
                        <div v-html="content" id="app">
                        </div>
                    </div>
                </div>
            </v-col>
            <v-col xl="4" lg="4" md="4" sm="12">
                <div class="my-10">
                    <v-card>
                        <v-card-title>
                            <h3>Lastest Posts</h3>
                        </v-card-title>
                        <v-card-text v-for="(item, index) in lastestPosts" :key="index">
                            <nuxt-link :to="{ path: '/postDetail?id=' + item.id }">
                                <h4 class="mt-2">{{ item.title }}</h4>
                            </nuxt-link>
                        </v-card-text>
                    </v-card>
                </div>
                <div class="my-10">
                    <v-card>
                        <v-card-title>
                            <h3>Categorys</h3>
                        </v-card-title>
                        <v-card-text v-for="(item, index) in tags" :key="index">
                            <nuxt-link :to="{ path: '/blog?search=' + item.name }">
                                <h4 class="mt-2">{{ item.name }}</h4>
                            </nuxt-link>
                        </v-card-text>
                    </v-card>
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
        category: String,
        content: String,
        image: String,
        lastestPosts: Array,
        tags: Array,
    },
};
</script>

<style >
</style>
