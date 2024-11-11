<script>
import Photo from '@/components/Photo.vue';
const token = sessionStorage.getItem('token');

export default {
    mounted() {
        if (localStorage.getItem('reloadedstream')) {
            localStorage.removeItem('reloadedstream');
        } else {
            localStorage.setItem('reloadedstream', '1');
            location.reload();
        }
    },
    data() {
        return {
            userName: '',
            found: false,
            followCount: 0,
            followedCount: 0,
            photoCount: 0,
            isBanned: false,
            isFollowed: false,
            isItMe: false,
            photoList: [],
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
        const uid= this.$route.params.uid;
        const token = sessionStorage.getItem('token');
        this.isItMe = (uid == token);
        this.fetchUserData();
    },
    methods: {
        refresh() {
            location.reload();
        },
        async fetchUserData() {
            const uid = this.$route.params.uid;
            const token = sessionStorage.getItem('token');
            try {
                const response = await this.$axios.get(`/users/${uid}/profile`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                this.found = true;
                this.userName = response.data.Username;
                this.followCount = response.data.FollowCount;
                this.followedCount = response.data.FollowedCount;
                this.photoCount = response.data.PhotoCount;
                this.isBanned = response.data.IsBanned;
                this.isFollowed = response.data.IsFollowed;
                this.photoList = response.data.Photolist;

            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 400:
                            console.error('Bad request');
                            this.userName = "You have to login first"
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            this.userName = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            this.userName = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            if (uid === "null") {
                                this.userName = "You are not logged in";
                            }
                            else {
                                this.userName = "User not found";
                            }
                            break;
                        case 500:
                            console.error('Internal Server Error:', error.response.data);
                            this.userName = "Internal Server Error"
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error.message);
                }
            }
        },
        async toggleFollow() {
            // frontend
            this.isFollowed = !this.isFollowed;
            // backend
            const uid = this.$route.params.uid;
            const token = sessionStorage.getItem('token');
            try {
                if (this.isFollowed) {
                    this.followCount += 1;
                    await this.$axios.put(`/users/${token}/follows/${uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    this.followCount -= 1;
                    await this.$axios.delete(`/users/${token}/follows/${uid}`, {
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
            // frontend
            this.isBanned = !this.isBanned;
            // backend
            const uid = this.$route.params.uid;
            const token = sessionStorage.getItem('token');
            try {
                if (this.isBanned) {
                    await this.$axios.put(`/users/${token}/bans/${uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    await this.$axios.delete(`/users/${token}/bans/${uid}`, {
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
        Photo,
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
        <div class="photos">
            <Photo v-for="photo in photoList" :key="photo.Pid" :pid="photo.Pid" :Date="photo.Date"
                :authorName="userName" :likeCount="photo.likecount" :caption="photo.caption" />
        </div>
    </div>
</template>