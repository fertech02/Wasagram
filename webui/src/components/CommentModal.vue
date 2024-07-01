<script>
export default {
    props: ['log', 'token'],
    data() {
        return {}
    },

    methods: {
        async uncommentPhoto(uid, pid) {
            try {
                let response = await this.$axios.delete('/photos/'+pid+'/comments/'+uid, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': "Bearer " + localStorage.getItem('token')       
                    }
                })
                location.reload();
                if (response.ok) {
                    this.$emit('uncomment', this.log.commentId);
                } else {
                    console.log('Error deleting comment');
                }
            } catch (error) {
                console.log('Error deleting comment');
            }
        },
    },
}
</script>

<template>
    <div>
        <b-modal id="commentModal" title="Comment" hide-footer>
            <div class="d-flex justify-content-between">
                <p>{{ log.comment }}</p>
                <button @click="uncommentPhoto(log.commentId, log.photoId)" class="btn btn-danger">Delete</button>
            </div>
        </b-modal>
    </div>
</template>

<style scoped>
.modal textarea {
    font-family: 'Roboto', sans-serif;
    width: 100%;
    height: 100px;
}
</style>