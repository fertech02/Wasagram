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
    Pid: String,
    Uid: String,
    Date: String,
  },

  data() {
    return {
      imgSrc: null,
      isLiked: false,
      LikeCount: 0,
      Author: '',
      isMe: false,
      notBanned: true,
      modalId: this.Pid,
    };
    
  },

  async created() {
    if (this.Pid) {
      try {
        const response = await this.$axios.get(`/photos/${this.Pid}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          responseType: 'blob',
        });
        const imageUrl = URL.createObjectURL(response.data);
        this.imgSrc = imageUrl;

        const isL = await this.$axios.get(`/photos/${this.Pid}/likes/${token}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.isLiked = isL.data.isLiked;

        const LikeCount = await this.$axios.get(`/photos/${this.Pid}/likes`, {
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
        const response = await this.$axios.delete(`/photos/${this.Pid}`, {
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
          await this.$axios.put(`/photos/${this.Pid}/likes/${token}`, {
          }, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
        } else {
          this.LikeCount -= 1;
          await this.$axios.delete(`/photos/${this.Pid}/likes/${token}`, {
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
          <div class="author">{{ authorName }}, {{ date }}</div>
          <div class="actions">
            <button @click="likePhoto" class="btn btn-sm btn-outline-primary ms-3">
              {{ isLiked ? 'Unlike' : 'Like' }}
            </button>
            <span class="like-counter">{{ LikeCount }} Likes <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#thumbs-up" />
              </svg></span>
            <button @click="commentPhoto" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal"
              :data-bs-target="'#usersModal' + modalId">
              Comment <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#message-circle" />
              </svg>
            </button>
            <button @click="viewComments" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal"
              :data-bs-target="'#listModal' + modalId">
              View Comments <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#message-square" />
              </svg>
            </button>
            <CommentModal :Pid="this.modalId" />
            <CommentListModal :Pid="this.modalId" />
          </div>
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