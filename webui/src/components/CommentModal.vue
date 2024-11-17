<script>
const token = sessionStorage.getItem('token');

export default {
    props: {
        pid: String,
        uid: String,
    },

    data() {
        return {
            commentPostTry: false,
            commentText: '',
            Text: '',
        };
    },

    methods: {
        async commentPhoto() {
            console.log("Posting Comment: ", this.commentText);
            this.commentPostTry = true;
            try {
                const config = {
                    headers: {
                        'Authorization': `Bearer ${token}`,
                    },
                };
                const response = await this.$axios.post(`/photos/${this.pid}/comments/${this.uid}`, { Message: this.commentText }, config);
                this.Text = "Comment Posted!";
                location.reload();
            }
            catch {
                console.error(error.response.data);
                this.Text = "Error posting comment";
            }
        },
    },
};

</script>


<template>
    <div class="modal fade" :id="'usersModal' + pid" tabindex="-1" aria-labelledby="ModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" id="ModalLabel">Post a comment for photo</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <form @submit.prevent="commentPhoto">
                        <div class="mb-3">
                            <input v-model="commentText" class="form-control" id="commentText" rows="4"
                                placeholder="Enter your comment" />
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                            <button type="submit" class="btn btn-primary">Post Comment</button>
                        </div>
                    </form>
                    <p v-if="commentPostTry" class="mt-3" style="font-size: 25px;">{{ Text }}</p>
                </div>
            </div>
        </div>
    </div>
</template>