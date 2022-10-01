<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-10">
            <v-col xl="6" lg="6" md="6" sm="8" xs="12">
                <nuxt-link to="/admin/category">
                    <v-btn color="secondary" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Edit Catégorie</v-card-title>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field @keydown.enter.prevent="saveBtn" label="Saisir une Catégorie" class="form__control" v-model="categoryField" hide-details="auto">
                            </v-text-field>

                            <v-btn @click="saveBtn" color="error" class="login__btn">Sauvegarder</v-btn>
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
    addTag,
    updateTag,
    getAllTags
} from '/functions/tag.js'

export default {
    name: "EditCategoryPage",
    head() {
        return {
            title: 'Update-Category',
        }
    },
    components: {
        SidebarVue,
    },
    data() {
        return {
            id: null,
            categoryField: '',
        }
    },
    methods: {
        saveBtn() {
            if (this.categoryField.trim().length === 0) {
                this.$toast.warning("Veuillez saisir la catégorie !");
            } else {
                getAllTags().then(result => {
                    var found = false

                    for (const item in result) {
                        if (result[item].name === this.categoryField) {
                            this.$toast.warning("La nouvelle catégorie existe déjà !");
                            found = true
                            break
                        }
                    }

                    if (found == false) {
                        if (this.id == null) {
                            addTag(this.$auth.strategy.token.get(), this.categoryField)
                        } else {
                            updateTag(this.$auth.strategy.token.get(), parseInt(this.id), this.categoryField)
                        }

                        this.$router.push({
                            path: '/admin/category'
                        })
                    }
                });
            }
        }
    },
    mounted: function () {
        if (this.$route.query.id != null) {
            this.id = parseInt(this.$route.query.id)
            this.categoryField = this.$route.query.name
        }
    },
};
</script>

<style>
</style>
