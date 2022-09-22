<template>
<div>
    <SidebarVue />
    <v-container>
        <v-card>
            <v-card-text class="d-flex justify-space-between align-center">
                <v-card-title v-if="tags != null"> Catégories ( {{ tags.length }} ) </v-card-title>
                <v-card-actions>
                    <nuxt-link to="/admin/create-category">
                        <v-btn color="red lighten-2">New </v-btn>
                    </nuxt-link>
                </v-card-actions>
            </v-card-text>
            <hr />

            <v-card-subtitle>
                <v-simple-table>
                    <template v-slot:default>
                        <thead>
                            <tr>
                                <th class="text-left">Catégories</th>
                                <th class="text-left">Editer</th>
                                <th class="text-left">Supprimer</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in tags" :key="index">
                                <td class="text-left">{{ item.id }}</td>
                                <td class="text-left">
                                    <nuxt-link to="/admin/update-category/1">
                                        <v-btn text class="ma-2" @click="edit_dialog = true" variant="text" color="blue">
                                            <v-icon color="blue darken-3"> mdi-pencil</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                                <td class="text-left">
                                    <v-btn text class="ma-2" variant="text" color="red">
                                        <v-icon color="red darken-3"> mdi-delete</v-icon>
                                    </v-btn>
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
import SidebarVue from "../../components/admin/Sidebar.vue";

import {
    getAllTags
} from '../../functions/post.js'

export default {
    name: "CategoryPage",
    head() {
        return {
            title: 'Category',
        }
    },
    data: () => ({
        tags: null
    }),
    components: {
        SidebarVue,
    },
    methods: {
        getAllTags
    },
    mounted: function () {
        this.getAllTags().then(result => {
            if (result != null) {
                this['tags'] = result
            }
        });
    },
};
</script>

<style>
</style>
