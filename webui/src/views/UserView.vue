<script>
import CommentModal from '../components/CommentModal.vue';

export default {

    mounted() {
        if (localStorage.getItem("reloadedStream")) {
            localStorage.removeItem("reloadedStream");
        } else {
            localStorage.setItem("reloadedStream", "1");
            location.reload();
        }
    },
    components: {
        CommentModal,
    },
    data: function() {
        return {
            
            // Profile
            Profile: {},

            // Photos
            Photos: {},

            // Follow
            Follow: {},

            // Ban
            Ban: {},

            // Comments
            Comments: {
                Uid: '',
                Pid: '',
                comments: [
                    {
                        Uid: '',
                        Pid: '',
                        comment: '',
                    }
                ],
            },
            comment: '',
        }
    },
    watch: {
        '$route.params.Uid' (newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        }
    },
    async created() {
        const Uid = this.$route.params.Uid;
        // this.fetchUserData();
    },
    methods: {
        async refresh(){},
        async getUserProfile(){},
        async getUserPhotos(){},
        async banUser(){},
        async unbanUser(){},
        async followUser(){},
        async unfollowUser(){},
        async commentPhoto(){},
        async openLog(){},
        async likePhoto(){},
        async unlikePhoto(){},
    }

};
</script>

<template>
</template>