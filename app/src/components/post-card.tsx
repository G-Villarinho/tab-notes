import { useState } from "react";
import { formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Heart, AlertTriangle } from "lucide-react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { clsx } from "clsx";

import { Button } from "@/components/ui/button";
import { UserAvatar } from "@/components/user-avatar";
import { likePost } from "@/api/like-post";
import { unlikePost } from "@/api/unlike-post";
import { GetFeedResponse } from "@/api/get-feed";
import { Link } from "react-router-dom";

interface PostCardProps {
  post: {
    postId: string;
    title: string;
    content: string;
    likes: number;
    authorName: string;
    authorUsername: string;
    createdAt: string;
    likedByUser: boolean;
  };
}

export function PostCard({ post }: PostCardProps) {
  const [liked, setLiked] = useState(post.likedByUser);
  const [likes, setLikes] = useState(post.likes);
  const [animating, setAnimating] = useState(false);
  const queryClient = useQueryClient();

  const likeMutation = useMutation({
    mutationFn: likePost,
    onSuccess: () => {
      updateFeedLike(true);
    },
  });

  const unlikeMutation = useMutation({
    mutationFn: unlikePost,
    onSuccess: () => {
      updateFeedLike(false);
    },
  });

  function handleLikeToggle() {
    setAnimating(true);
    setTimeout(() => setAnimating(false), 200);

    if (liked) {
      setLiked(false);
      setLikes((prev) => prev - 1);
      unlikeMutation.mutate({ postId: post.postId });
    } else {
      setLiked(true);
      setLikes((prev) => prev + 1);
      likeMutation.mutate({ postId: post.postId });
    }
  }

  function updateFeedLike(liked: boolean) {
    queryClient.setQueryData<GetFeedResponse[]>(["feed"], (old) =>
      old?.map((p) =>
        p.postId === post.postId ? { ...p, likedByUser: liked } : p
      )
    );
  }

  const timeAgo = formatDistanceToNow(post.createdAt, {
    addSuffix: true,
    locale: ptBR,
  });

  return (
    <div className="flex gap-4">
      <div className="flex flex-col items-center pt-1">
        <Link to={`/${post.authorUsername}`}>
          <UserAvatar
            name={post.authorName}
            username={post.authorUsername}
            className="h-10 w-10"
          />
        </Link>
        <div className="w-px flex-1 bg-muted mt-2" />
      </div>

      <div className="flex-1 rounded-xl border p-4 shadow-sm bg-card hover:shadow-md transition space-y-2 relative">
        <div className="flex justify-between items-start gap-2">
          <div className="flex flex-col">
            <Link
              to={`/${post.authorUsername}`}
              className="font-medium text-sm hover:underline transition"
              title={post.authorName}
            >
              @{post.authorUsername}
            </Link>
            <span className="text-xs text-muted-foreground">{timeAgo}</span>
          </div>

          <Button variant="outline" size="xs">
            <AlertTriangle className="w-2 h-2 text-zinc-500" />
          </Button>
        </div>

        <div>
          <h3 className="font-semibold text-base">{post.title}</h3>
          <p className="text-sm text-muted-foreground mt-1">{post.content}</p>
        </div>

        <div className="pt-2 flex items-center gap-3">
          <Button
            onClick={handleLikeToggle}
            variant="outline"
            size="sm"
            className={clsx(
              "gap-1 text-sm transition-transform duration-200",
              liked && "text-pink-500 bg-pink-100 border-pink-200",
              animating && "scale-105"
            )}
          >
            <Heart className="w-4 h-4 fill-current" />
            Curtir
          </Button>

          <span className="text-sm text-muted-foreground">
            {likes} {likes === 1 ? "curtida" : "curtidas"}
          </span>
        </div>
      </div>
    </div>
  );
}
