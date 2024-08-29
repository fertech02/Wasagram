<script>
  import CommentModal from '@/components/CommentModal.vue';
  import CommentListModal from '@/components/CommentListModal.vue';
  
  const token = sessionStorage.getItem('authToken');
  export default {
    components: {
      CommentModal,
      CommentListModal,
    },
    props: {
      Pid: String,
      Uid: String,
      File: String,
      date: String,
    },
    data() {
      return {
        imgSrc: this.File,
        authorId: this.Uid,
        isLiked: false,
        isMe: false,
        notBanned: true,
        caption: '',
        LikeCount: 0,
      };
    },
  
    async created() {
      if (this.Pid) {
        try {
          const response = await this.$axios.get('/photos/' + this.Pid, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
        const imageUrl = URL.createObjectURL(response.data);
        this.imgSrc = imageUrl;
        const isL = await this.$axios.get(`/photos/${this.Pid}/likes/${token}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.isLiked = isL.data.isLiked
        this.findAuthorId();
        }
        catch (error) {
          console.error(error);
        }
      } else {
        console.error('Error:', error)
      }
    },
    computed: {
  
    },
    methods: {

      async findAuthor() {
        try {
          const response = this.$route.params.uid;
          const hasStream = this.$route.path.includes('/stream');
          if (userId == token && !hasStream) {
            this.isMe = true;
          }

        } catch (error) {
          console.error(error);
        }
      },

      async deletePhoto(){
        try {
          const response = await this.$axios.delete(`/photos/${this.Pid}`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          location.reload();
        } catch (error) {
          console.error(error);
        }
      },

      async likePhoto() {
        this.isLiked = !this.isLiked;
        try {
          if (this.isLiked) {
            this.LikeCount += 1;
            const response = await this.$axios.put('/photos/' + this.Pid + '/likes' + this.Uid, {}, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          } else {
            this.LikeCount -= 1;
            await this.$axios.delete('/photos/' + this.Pid + '/likes' + this.Uid, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            });
          }

        } catch (error) {
          console.error(error);
        }
      },
    }
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
          <div class="card-text text-center bg-light fs-5">{{ caption }}</div>
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
            <button @click="getComments" class="btn btn-sm btn-outline-secondary" data-bs-toggle="modal"
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
  
  .caption {
    display: flex;
    align-items: center;
    margin-top: 10px;
  }
  
  .caption-border {
    flex: 1;
    height: 3px;
    background-color: #1a1212;
    padding: 4px;
    margin-top: 10px;
    margin-bottom: 10px;
  
  }
  
  .caption-text {
    padding: 0 10px;
  }
  </style>