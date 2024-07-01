<script>
import CommentModal from '../components/CommentModal.vue'

export default {
    components: {
        CommentModal
    },
    data: function() {
        return {
            errorMsg: null,
            // Get User Profile
            Profile : {
                id : '',
                username: '',
                photoCount: 0,
                followerCount: 0,
                followingCount: 0,
                photos: [],
            },

            // Get User Photos
            Photos : [],

            // Get FollowedUsers
            FollowedUsers : [],
            
            // Get Followees
            Followees : [],
        }
    },
    methods: {
        async refresh() {
            await this.getUserProfile(),
            await this.getUser()
        },

        async getUserProfile(){},
        async getUser(){},
        async deletePhoto(){},
        async setMyUserName(){},
        async commentPhoto(){},
        async openLog(){},
        async likePhoto(){},
        async unlikePhoto(){},
        async getFolowedUsers(){},
        async getFollowees(){},
    }, 
}
</script>

<template>
</template>
