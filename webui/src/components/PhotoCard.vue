<script>
import CommentModal from '@/components/CommentModal.vue';
import CommentListModal from '@/components/CommentListModal.vue';

const token = sessionStorage.getItem('token');

export default {

  components: {
    CommentModal,
    CommentListModal,
  },

  props: {
    pid: String,
    uid: String,
    date: String,
  },

  data() {
    return {
      imgSrc: null,
      isLiked: false,
      LikeCount: 0,
      Author: '',
      isMe: false,
      notBanned: true,
      modalId: this.pid,
    };
    
  },

  async created() {
    if (this.Pid) {
      try {
        const response = await this.$axios.get(`/photos/${this.pid}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          responseType: 'blob',
        });
        const imageUrl = URL.createObjectURL(response.data);
        console.log(imageUrl);
        this.imgSrc = imageUrl;

        const isL = await this.$axios.get(`/photos/${this.pid}/likes/${token}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.isLiked = isL.data.isLiked;

        const LikeCount = await this.$axios.get(`/photos/${this.pid}/likes`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.LikeCount = LikeCount.data
  
        this.findAuthorId();
      } 
      catch (error) {
        if (error.response) {
          const statusCode = error.response.status;
          this.notBanned = false;
          switch (statusCode) {
            case 401:
              console.error('Access Unauthorized:', error.response.data);
              break;
            case 403:
              console.error('Access Forbidden:', error.response.data);
              break;
            case 404:
              console.error('Not Found:', error.response.data);
              break;
            default:
              console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
          }
        } else {
          console.error('Error:', error);
        }
      }
    }
  },

  computed: {

  },

  methods: {

    async findAuthorId() {
      try {
        const uid = this.$route.params.uid;
        const hasStreamSegment = this.$route.path.includes('/stream');
        if (uid == token && !hasStreamSegment) {
          this.isMe = true;
        };
      }
      catch (error) {
        console.error(error, "Error searching photo owner.")
      }
    },

    async deletePhoto() {
      try {
        const response = await this.$axios.delete(`/photos/${this.pid}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          }
        },);
        location.reload();
      }
      catch (error) {
        console.error(error, "cant delete photo!")
      }
    },

    async likePhoto() {

      this.isLiked = !this.isLiked;
      
      try {
        const token = sessionStorage.getItem('authToken');
        if (this.isLiked) {
          this.LikeCount += 1;
          await this.$axios.put(`/photos/${this.pid}/likes/${token}`, {
          }, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
        } else {
          this.LikeCount -= 1;
          await this.$axios.delete(`/photos/${this.pid}/likes/${token}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });

        }
      } catch (error) {
        console.error(error, "Error modifying like status.")
      }

    },
  },
  
};
</script>

<template>
  <div class="center-container">
    <div class="photo-card">
      <img :src="imgSrc" alt="Photo" width="400" height="400" />
      <div class="photo-details">
        <div class="author">
          <span>Author: {{ Author }}</span>
        </div>
        <div class="actions">
          <div>
            <button v-if="isMe" @click="deletePhoto" class="btn btn-danger">Delete</button>
            <button @click="likePhoto" class="btn btn-primary">{{ isLiked ? 'Unlike' : 'Like' }}</button>
            <span class="like-counter">{{ LikeCount }}</span>
          </div>
          <CommentModal :pid="pid" :uid="uid" />
          <CommentListModal :pid="pid" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.center-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.photo-card {
  border: 3px solid #6d6969;
  border-radius: 4px;
  padding: 10px;
  width: 500px;
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.photo-details {
  margin-top: 10px;
}

.author {
  font-size: 20px;
  margin-bottom: 5px;
}

.actions {
  display: flex;
  justify-content: space-between;
  margin: 15px;
}

.like-counter {
  margin-left: 2px;
  border: 2px solid #d102027a;
  border-radius: 4px;
  padding: 8px;
}
</style>