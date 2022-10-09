<template>
<div>
    <SidebarVue />
    <v-container>
        <v-card>
            <v-card-text class="d-flex justify-space-between align-center">
                <v-card-title v-if="comments != null"> Comments ({{ comments.length }}) </v-card-title>
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
                                <th class="text-enter">Username</th>
                                <th class="text-enter">Comment</th>
                                <th class="text-enter">Post</th>
                                <th class="text-enter">View</th>
                                <th class="text-enter">Valid</th>
                                <th class="text-enter">Delete</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in comments" :key="index">
                                <td class="">{{ item.id }}</td>
                                <td class="">{{ item.username }}</td>
                                <td class="">{{ item.message }}</td>
                                <td class="">{{ item.message }}</td>
                                <td class="">
                                    <nuxt-link :to="'/postDetail?id=' + item.postId">
                                        <v-btn text class="ma-2" @click="view_dialog = true" variant="text" color="blue">
                                            <v-icon color="green darken-3">mdi-eye</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                                <td class="">
                                    <v-switch @change="validityBtn(item.id, item.valid)" v-model="item.valid"></v-switch>
                                </td>
                                <td class="">
                                    <nuxt-link :to="{ path: '/admin/commentAction/removeComment?id=' + item.id + '&name=' + item.title }">
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
    getAllMessages,
    updateValidityMessage
} from '/functions/message.js'

export default {
    name: "CommentPage",
    head() {
        return {
            title: 'Comment'
        }
    },
    components: {
        SidebarVue,
    },
    data: () => ({
        comments: null,
    }),
    methods: {
        validityBtn(id, valid) {
            updateValidityMessage(this.$auth.strategy.token.get(), id, valid)
            this.$router.push({path: '/admin/comment'})
        }
    },
    mounted: function () {
        getAllMessages(this.$auth.strategy.token.get()).then(result => {
            if (result != null) {
                this.comments = result
            }
        });
    },
};
</script>

<style>
</style>
