<template>
<div>
    <SidebarVue />
    <v-container>
        <v-card>
            <v-card-text class="d-flex justify-space-between align-center">
                <v-card-title v-if="location != null"> Locations ({{ location.length }}) </v-card-title>
                <v-card-actions>
                    <nuxt-link to="/admin/locationAction/editLocation?id=0">
                        <v-btn color="red lighten-2">New</v-btn>
                    </nuxt-link>
                </v-card-actions>
            </v-card-text>
            <hr />

            <v-card-subtitle>
                <v-simple-table>
                    <template v-slot:default>
                        <thead>
                            <tr>
                                <th class="text-left">Nom</th>
                                <th class="text-left">Image</th>
                                <th class="text-left">Latitude</th>
                                <th class="text-left">Longitude</th>
                                <th class="text-left">Editer</th>
                                <th class="text-left">Supprimer</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in location" :key="index">
                                <td class="text-left">{{ item.name }}</td>
                                <td class="">
                                    <v-img :src="item.image" width="30" height="30" cover></v-img>
                                </td>
                                <td class="text-left">{{ item.latitude }}</td>
                                <td class="text-left">{{ item.longitude }}</td>
                                <td class="text-left">
                                    <nuxt-link :to="{ path: '/admin/locationAction/editLocation?id=' + item.id }">
                                        <v-btn text class="ma-2" @click="edit_dialog = true" variant="text" color="blue">
                                            <v-icon color="blue darken-3">mdi-pencil</v-icon>
                                        </v-btn>
                                    </nuxt-link>
                                </td>
                                <td class="text-left">
                                    <nuxt-link :to="{ path: '/admin/locationAction/removeLocation?id=' + item.id + '&name=' + item.name}">
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
    <MapVue />
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";
import MapVue from "/components/front/Map.vue";

import {
    getAllLocations,
} from '/functions/location.js'

export default {
    name: "LocationPage",
    head() {
        return {
            title: 'Location',
        }
    },
    data: () => ({
        location: null
    }),
    components: {
        SidebarVue,
        MapVue,
    },
    mounted: function () {
        getAllLocations().then(result => {
            if (result != null) {
                this.location = result
            }
        });
    },
};
</script>

<style>
</style>
