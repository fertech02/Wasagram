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
    photoId: String,
    userId: String,
    date: String,
  },

  data() {
    return {
      imgSrc: null,
      isLiked: false,
      LikeCount: 0,
      Author: this.userId,
      isMe: false,
      notBanned: true,
      Pid: this.photoId,
    };
    
  },

  async created() {
    if (this.photoId) {
      try {
        console.log(this.photoId);
        const response = await this.$axios.get(`/photos/${this.photoId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          responseType: 'blob',
        });
        const imageUrl = URL.createObjectURL(response.data);
        this.imgSrc = imageUrl;

        const isL = await this.$axios.get(`/photos/${this.photoId}/likes/${token}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.isLiked = isL.data.is_liked;
        console.log(this.isLiked);

        const LikeCount = await this.$axios.get(`/photos/${this.photoId}/likes`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.LikeCount = LikeCount.data.like_count;
        console.log(this.LikeCount);
  
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
        const response = await this.$axios.delete(`/photos/${this.photoId}`, {
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
        if (this.isLiked) {
          this.LikeCount += 1;
          console.log(this.photoId, token);
          await this.$axios.put(`/photos/${this.photoId}/likes/${token}`, {
          }, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
        } else {
          this.LikeCount -= 1;
          console.log(this.photoId, token);
          await this.$axios.delete(`/photos/${this.photoId}/likes/${token}`, {
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
  <div class="container mt-5" v-if="notBanned">
    <div class="center-container">
      <div class="card photo-card">
        <button v-if="isMe" @click="deletePhoto" class="btn btn-danger delete-button mb-2">
          Delete Photo <svg class="feather">
            <use href="/feather-sprite-v4.29.0.svg#trash-2" />
          </svg>
        </button>

        <img :src="imgSrc" alt="Photo" class="card-img-top" />
        <div class="card-body photo-details">
          <div class="author">{{ Author }}, {{ this.date }}</div>
          <div class="actions">
            <button @click="likePhoto" class="btn btn-sm btn-outline-primary ms-3">
              {{ isLiked ? 'Unlike' : 'Like' }}
            </button>
            <span class="like-counter">{{ LikeCount }} Likes <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#thumbs-up" />
              </svg></span>
          </div>
            <button @click="commentPhoto" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal"
              :data-bs-target="'#usersModal' + this.photoId">
              Comment <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#message-circle" />
              </svg>
            </button>
            <button @click="viewComments" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal"
              :data-bs-target="'#listModal' + this.photoId">
              View Comments <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#message-square" />
              </svg>
            </button>
            <CommentModal :photoId="Pid"/>
            <CommentListModal :photoId="Pid"/>
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