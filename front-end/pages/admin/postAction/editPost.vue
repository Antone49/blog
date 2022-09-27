<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-12">
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <nuxt-link to="/admin/post">
                    <v-btn color="secondary" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Edite Post</v-card-title>
                    <br>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field label="Titre" v-model="title" class="form__control" hide-details="auto" outlined>
                            </v-text-field>
                            <v-select class="mt-3" label="Catégories" v-model="tagsPost" :items="allTags" multiple outlined>
                                <template v-slot:selection="{ item }">
                                    <v-chip>
                                        {{item.name}}
                                    </v-chip>
                                </template>
                                <template v-slot:item="{item}">
                                    <v-icon v-if="tagsPost !== null && tagsPost.includes(item)" color="primary" class="mr-3">
                                        mdi-checkbox-marked
                                    </v-icon>
                                    <v-icon v-else class="mr-3">
                                        mdi-checkbox-blank-outline
                                    </v-icon>
                                    {{item.name}}
                                </template>
                            </v-select>

                            <v-select class="mt-3" label="Locations" v-model="locationsPost" :items="allLocations" multiple outlined>
                                <template v-slot:selection="{ item, index }">
                                    <v-chip>
                                        <img :src="'/images/location/' + item.image" width="25" height="25">
                                        {{item.name}}
                                    </v-chip>
                                </template>
                                <template v-slot:item="{item}">
                                    <v-icon v-if="locationsPost !== null && locationsPost.includes(item)" color="primary" class="mr-3">
                                        mdi-checkbox-marked
                                    </v-icon>
                                    <v-icon v-else class="mr-3">
                                        mdi-checkbox-blank-outline
                                    </v-icon>
                                    <img :src="'/images/location/' + item.image" width="30" height="30">
                                    {{item.name}}
                                </template>
                            </v-select>

                            <v-file-input accept="image/*" label="Image" outlined></v-file-input>
                            <vue-editor v-model="content" />

                            <v-btn color="error" @click="saveBtn" class="login__btn">Sauvegarder</v-btn>
                        </form>
                    </v-card-subtitle>
                </v-card>
            </v-col>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <PostDetailVue :category="category" :title="title" :content="content" :image="image" />
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";
import PostDetailVue from "/components/front/PostDetail.vue";

import {
    addPost,
    updatePost,
    getPost,
    getPostTags,
    updatePostTags,
    getPostLocations,
    updatePostLocations
} from '/functions/post.js'

import {
    getAllLocations
} from '/functions/location.js'

import {
    getAllTags
} from '/functions/tag.js'

export default {
    name: "editPostPage",
    head() {
        return {
            title: 'EditPost',
        }
    },
    components: {
        SidebarVue,
        PostDetailVue,
    },
    data: () => ({
        id: 0,
        title: null,
        content: null,
        category: null,
        image: "loading.png",
        allTags: [],
        tagsPost: [],
        allLocations: [],
        locationsPost: [],
    }),
    mounted: function () {
        this.id = parseInt(this.$route.query.id)

        getAllTags().then(result => {
            if (result != null) {
                this.allTags = result

                if (this.id != 0) {
                    getPostTags(this.id).then(result => {
                        if (result != null) {
                            for (const i in result) {
                                for (const j in this.allTags) {
                                    if (result[i].id == this.allTags[j].id) {
                                        this.tagsPost.push(this.allTags[j])
                                    }
                                }
                            }
                        }
                    });
                }
            }
        });

        getAllLocations().then(result => {
            if (result != null) {
                this.allLocations = result

                if (this.id != 0) {
                    getPostLocations(this.id).then(result => {
                        if (result != null) {
                            for (const i in result) {
                                for (const j in this.allLocations) {
                                    if (result[i].id == this.allLocations[j].id) {
                                        this.locationsPost.push(this.allLocations[j])
                                    }
                                }
                            }
                        }
                    });
                }
            }
        });

        if (this.id != 0) {
            getPost(this.id).then(result => {
                if (result != null) {
                    this.title = result.title
                    this.content = result.content
                    this.image = result.image
                }
            });
        }
    },
    methods: {
        saveBtn() {
            if (this.id == 0) {
                addPost(this.$auth.strategy.token.get(), this.title, this.image, this.content)
            } else {
                updatePost(this.$auth.strategy.token.get(), this.id, this.title, this.image, this.content)

                var tagsId = []
                for (const i in this.tagsPost) {
                    tagsId.push(this.tagsPost[i].id)
                }

                updatePostTags(this.$auth.strategy.token.get(), this.id, tagsId)

                var locationId = []
                for (const i in this.locationsPost) {
                    locationId.push(this.locationsPost[i].id)
                }

                updatePostLocations(this.$auth.strategy.token.get(), this.id, locationId)
            }

            this.$router.push({
                path: '/admin/post'
            })
        }
    }
};
</script>

<style>
</style>
