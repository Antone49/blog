<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-10">
            <v-col xl="6" lg="6" md="6" sm="8" xs="12">
                <nuxt-link to="/admin/comment">
                    <v-btn color="secondary" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Supprime Commentaire</v-card-title>
                    <v-card-subtitle>
                        <form action="" class="form">
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
    removeMessage
} from '/functions/message.js'

export default {
    name: "RemoveCommentPage",
    head() {
        return {
            title: 'RemoveComment',
        }
    },
    components: {
        SidebarVue,
    },
    data() {
        return {
            id: '',
        }
    },
    methods: {
        removeBtn() {
            removeMessage(this.$auth.strategy.token.get(), this.id)
            this.$router.push({path: '/admin/comment'})
        }
    },
    mounted: function () {
        this.id = parseInt(this.$route.query.id)
    },
};
</script>

<style>
</style>
