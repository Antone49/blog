<template>
<div>
    <v-container>
        <v-row>
            <v-col xl="4" lg="4" md="4" sm="6">
                <v-card class="d-flex justify-space-between">
                    <v-card>
                        <v-card-title>Post</v-card-title>
                        <v-card-subtitle v-if="posts != null">{{ posts.length }}</v-card-subtitle>
                    </v-card>
                    <v-card-actions>
                        <v-icon large color="blue-grey darken-2">
                            mdi-post
                        </v-icon>
                    </v-card-actions>
                </v-card>
            </v-col>
            <v-col xl="4" lg="4" md="4" sm="6">
                <v-card class="d-flex justify-space-between">
                    <v-card>
                        <v-card-title>Category</v-card-title>
                        <v-card-subtitle>33</v-card-subtitle>
                    </v-card>
                    <v-card-actions>
                        <v-icon large color="blue-grey darken-2">
                            mdi-widgets
                        </v-icon>
                    </v-card-actions>
                </v-card>
            </v-col>
            <v-col xl="4" lg="4" md="4" sm="6">
                <v-card class="d-flex justify-space-between">
                    <v-card>
                        <v-card-title>Comments</v-card-title>
                        <v-card-subtitle>33</v-card-subtitle>
                    </v-card>
                    <v-card-actions>
                        <v-icon large color="blue-grey darken-2">
                            mdi-gavel
                        </v-icon>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
        <DashboardTableVue :posts="posts"/>
    </v-container>
</div>
</template>

<script>
import DashboardTableVue from "./DashboardTable.vue";

export default {
    name: "DasboardComp",
    components: {
        DashboardTableVue,
    },
    props: {
        posts: undefined,
    },
};
</script>

<style>
</style>
