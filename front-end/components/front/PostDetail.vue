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
                        <p>
                            {{ content }}
                        </p>
                    </div>
                </div>
                <v-btn color="error" class="login__btn">J'aime</v-btn>
                <h3>Comments( 44 )</h3>
                <hr />
            </v-col>
            <v-col xl="4" lg="4" md="4" sm="12">
                <div class="my-10">
                    <v-card>
                        <v-card-title>
                            <h3>Lastest Posts</h3>
                        </v-card-title>
                        <v-card-text v-for="(item, index) in lastestPosts" :key="index">
                            <nuxt-link :to="{ path: '/post-detail/' + item.id }">
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
                            <nuxt-link :to="{ path: '/blog?search=' + item.id }">
                                <h4 class="mt-2">{{ item.id }}</h4>
                            </nuxt-link>
                        </v-card-text>
                    </v-card>
                </div>

            </v-col>
        </v-row>
        <v-row>
            <v-col xl="7" lg="7" md="7" sm="12">
                <div v-for="i in 6" :key="i">
                    <v-sheet>
                        <h3>Hilal ahmad</h3>
                        <p>comment</p>
                        <hr />
                    </v-sheet>
                </div>
            </v-col>
            <v-col xl="5" lg="5" md="5" sm="12" xs="12">
                <form action="" class="form">
                    <v-text-field label="Enter Your Name" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-text-field label="Enter Your Email" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-text-field label="Enter Your Comment" class="form__control" hide-details="auto">
                    </v-text-field>
                    <v-btn color="error" class="login__btn"> Comment</v-btn>
                </form>
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
