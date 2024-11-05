<script>
import PhotoCard from '@/components/Photo.vue';
const token = sessionStorage.getItem('token');

export default {
    mounted() {
        if (localStorage.getItem('reloadedstream')) {
            localStorage.removeItem('reloadedstream');
            console.log("Not reloading page");
        } else {
            localStorage.setItem('reloadedstream', '1');
            console.log("Reloading page");
            location.reload();
        }
    },

    data() {
        return {
            user_id: '',
            username: '',
            followCount: 0,
            followedCount: 0,
            photoCount: 0,
            found: false,
            reloadFlag: true,
        };
    },

    watch: {
        '$route.params.uid'(newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        },
    },

    async created() {
        const userId = this.$route.params.uid;
        console.log("User ID: ", userId);
        const token = sessionStorage.getItem('token');
        this.isItMe = (userId == token);
        console.log("Is it me? ", this.isItMe);
        this.fetchUserData(token);
    },

    methods: {

        refresh() {
            location.reload();
        },

        async fetchUserData(token) {
            try {
                const userId = this.$route.params.uid;

                console.log("Session token: ", token);
                console.log("User ID: ", userId);
                const response = await this.$axios.get(`/users/${userId}/profile`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                

                this.user_id = response.data.user_id;
                this.username = response.data.username;
                this.followers = response.data.followers;
                this.followees = response.data.followees;
                this.photo_count = response.data.photo_count;

            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 400:
                            console.error('Bad request');
                            this.username = "You have to login first"
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            this.username = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            this.username = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            if (user_id === "null") {
                                this.username = "You are not logged in";
                            }
                            else {
                                this.username = "User not found";
                            }
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },

        async toggleFollow() {
            this.isFollowed = !this.isFollowed;
            const userId = this.$route.params.uid;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isFollowed) {
                    this.followCount += 1;
                    await this.$axios.put(`/users/${token}/follows/${userId}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    this.followCount -= 1;
                    await this.$axios.delete(`/users/${token}/follows/${userId}`, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                }
            } catch (error) {
                console.error(error, "Error modifying follow status.")
            }

        },

        async toggleBan() {
            
            this.isBanned = !this.isBanned;
            
            const userId = this.$route.params.uid;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isBanned) {
                    await this.$axios.put(`/users/${token}/bans/${userId}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    await this.$axios.delete(`/users/${token}/bans/${userId}`, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });

                }
            } catch (error) {
                console.error(error, "Error modifying ban status.")
            }
        },
    },
    components: {
        PhotoCard,
    },
};
</script>

<template>
    <div class="container mt-5">
        <h1 class="display-4" style="font-size: 50px;">{{ username }}</h1>
        <div v-if="found">
            <div>
                <div v-if="!isItMe">
                    <div class="btn-group mt-1">
                        <button @click="toggleFollow" class="btn btn-warning">
                            {{ isFollowed ? 'Unfollow' : 'Follow' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                            </svg>
                        </button>
                        <button @click="toggleBan" class="btn btn-danger">
                            {{ isBanned ? 'Unban' : 'Ban' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#slash" />
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
            <div style="font-size: 20px;" class="container mt-2">
                <div class="row bg-light p-4 shadow-lg">
                    <div class="row">Followers: {{ followCount }}</div>
                    <div class="row">Followed: {{ followedCount }}</div>
                    <div class="row">Photos: {{ photoCount }}</div>
                </div>
            </div>

        </div>
        <hr />
        <div class="photos">
        </div>
    </div>
</template>