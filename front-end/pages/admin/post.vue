<template>
<div>
    <SidebarVue />
    <v-container>
        <v-card>
            <v-card-text class="d-flex justify-space-between align-center">
                <v-card-title v-if="posts != null"> Posts ({{ posts.length }}) </v-card-title>
                <v-card-actions>
                    <nuxt-link to="/admin/postAction/editPost">
                        <v-btn color="red lighten-2"> Nouveau </v-btn>
                    </nuxt-link>
                </v-card-actions>
            </v-card-text>
            <hr />

            <v-card-subtitle>
                <v-simple-table>
                    <template v-slot:default>
                        <thead>
                            <tr>
                                <th class="text-enter">Id</th>
                                <th class="text-enter">Title</th>
                                <th class="text-enter">Category</th>
                                <th class="text-enter">image</th>
                                <th class="text-enter">Views</th>
                                <th class="text-enter">View</th>
                                <th class="text-enter">Edit</th>
                                <th class="text-enter">Delete</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in posts" :key="index">
                                <td class="">{{ item.id }}</td>
                                <td class="">{{ item.title }}</td>
                                <td class="">Php</td>
                                <td class="">
                                    <v-img :src="'http://localhost:8800' + item.image" width="70" height="60" cover></v-img>
                                </td>
                                <td class="">44</td>
                                <td class="">
                                    <nuxt-link :to="'/postDetail?id=' + item.id">
                                        <v-btn text class="ma-2" @click="view_dialog = true" variant="text" color="blue">
                                            <v-icon color="green darken-3">mdi-eye</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                                <td class="">
                                    <nuxt-link :to="'/admin/postAction/editPost?id=' + item.id">
                                        <v-btn text class="ma-2" variant="text" color="blue">
                                            <v-icon color="blue darken-3"> mdi-pencil</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                                <td class="">
                                    <nuxt-link :to="{ path: '/admin/postAction/removePost?id=' + item.id + '&name=' + item.title }">
                                        <v-btn text class="ma-2" variant="text" color="red">
                                            <v-icon color="red darken-3">mdi-delete</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                            </tr>
                        </tbody>
                    </template>
                </v-simple-table>
            </v-card-subtitle>
        </v-card>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";

import {
    getAllPosts,
} from '/functions/post.js'

export default {
    name: "PostPage",
    head() {
        return {
            title: 'Post',
        }
    },
    components: {
        SidebarVue,
    },
    data: () => ({
        posts: null
    }),
    mounted: function () {
        getAllPosts().then(result => {
            if (result != null) {
                this.posts = result
            }
        });
    },
};
</script>

<style>
</style>
