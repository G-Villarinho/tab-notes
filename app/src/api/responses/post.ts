export interface PostResponse {
  id: string;
  title: string;
  content: string;
  likes: number;
  createdAt: string;
  likedByUser: boolean;
}
