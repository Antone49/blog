<template>
<div>
    <v-container class="my-10">
        <div class="d-flex align-center my-4">
            <v-text-field @keydown.enter.prevent="search" label="Search Here ..." class="form__control" hide-details="auto" v-model="searchField" outlined>
            </v-text-field>
            <v-btn @click="search" color="success" text>Search</v-btn>
        </div>
        <h1>All Post</h1>
        <v-row class="my-5">
            <v-col xl="4" lg="4" md="4" sm="6" xs="12" v-for="(item, index) in datas" :key="index">
                <v-card>
                    <v-img height="250" :src="'images/' + item.image"></v-img>
                    <v-card-title>{{ item.title }}</v-card-title>
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
        <div class="text-center">
            <v-pagination v-model="page" :length="4" circle></v-pagination>
        </div>
    </v-container>
</div>
</template>

<script>
import {
    getAllPosts
} from '../../functions/post.js'

export default {
    data() {
        return {
            page: 1,
            datas: null,
            searchField: null
        }
    },
    methods: {
        getAllPosts,
        search(event) {
            this.getAllPosts(this['searchField']).then(result => {
                this['datas'] = result
            });
        }
    },
    mounted: function () {
        if (this.$route.query.search != null) {
            this['searchField'] = this.$route.query.search
        }

        this.getAllPosts(this['searchField']).then(result => {
            this['datas'] = result
        });
    },
};
</script>

<style>
</style>
