<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-10">
            <v-col xl="6" lg="6" md="6" sm="8" xs="12">
                <nuxt-link to="/admin/post">
                    <v-btn color="secondary" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Supprime Post</v-card-title>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field :readonly="true" class="form__control" v-model="postField" hide-details="auto">
                            </v-text-field>

                            <v-btn @click="removeBtn" color="error" class="login__btn">Supprimer</v-btn>
                        </form>
                    </v-card-subtitle>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";

import {
    removePost
} from '/functions/post.js'

export default {
    name: "RemoveCategoryPage",
    head() {
        return {
            title: 'RemoveCategory',
        }
    },
    components: {
        SidebarVue,
    },
    data() {
        return {
            postField: '',
        }
    },
    methods: {
        removeBtn() {
            removePost(this.$auth.strategy.token.get(), parseInt(this.$route.query.id))
            this.$router.push({path: '/admin/post'})
        }
    },
    mounted: function () {
        this.postField = this.$route.query.name
    },
};
</script>

<style>
</style>
