<template>
<div>
    <v-container class="my-10">
        <h1>Latest Post</h1>
        <v-row class="my-5">
            <v-col xl="4" lg="4" md="4" sm="6" xs="12" v-for="(item, index) in datas" :key="index">
                <v-card>
                    <v-img height="250" :src="'http://localhost:8800/' + item.image"></v-img>
                    <v-card-title>{{ item.title }}</v-card-title>
                    <v-card-subtitle>category</v-card-subtitle>
                    <v-card-text>
                        <div v-html="item.content.substr(0, 50)" id="app">
                        </div>
                        <nuxt-link :to="{ path: '/postDetail?id=' + item.id }">
                            <v-btn class="mb-2 btn__border" variant="contained" color="primary" text>Read More</v-btn>
                        </nuxt-link>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import {
    getLastestPosts
} from '../../functions/post.js'

export default {
    data() {
        return {
            page: 1,
            datas: null,
        }
    },
    mounted: function () {
        getLastestPosts(3).then(result => {
            this.datas = result
        });
    },
};
</script>

<style>
</style>
