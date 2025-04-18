import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Loader2, Send } from "lucide-react";
import { UserAvatar } from "@/components/user-avatar";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormValidationError } from "@/components/form-validation-error";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createPost, CreatePostResponse } from "@/api/create-post";
import toast from "react-hot-toast";
import { GetFeedResponse } from "@/api/get-feed";

const createPostSchema = z.object({
  title: z.string().min(1, "Título é obrigatório."),
  content: z.string().min(1, "Conteúdo é obrigatório."),
});

type CreatePostSchema = z.infer<typeof createPostSchema>;

interface CreatePostCardProps {
  user: {
    name: string;
    username: string;
  };
}

export function CreatePostCard({ user }: CreatePostCardProps) {
  const queryClient = useQueryClient();

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<CreatePostSchema>({
    resolver: zodResolver(createPostSchema),
  });

  function onChangeFeedCache(newPost: CreatePostResponse) {
    const formatted: GetFeedResponse = {
      postId: newPost.id,
      title: newPost.title,
      content: newPost.content,
      likes: newPost.Likes,
      createdAt: newPost.createdAt,
      authorName: user.name,
      authorUsername: user.username,
      likedByUser: false,
    };

    queryClient.setQueryData<GetFeedResponse[]>(["feed"], (old) => [
      formatted,
      ...(old ?? []),
    ]);
  }

  const { mutateAsync: createPostFn, isPending } = useMutation({
    mutationFn: createPost,
    onSuccess: (newPost) => onChangeFeedCache(newPost),
  });

  async function handleCreatePost(data: CreatePostSchema) {
    try {
      await createPostFn(data);
      reset();
    } catch {
      toast.error("Erro ao criar o post");
    }
  }

  return (
    <div className="flex gap-4">
      <div className="flex flex-col items-center pt-1">
        <UserAvatar
          name={user.name}
          username={user.username}
          className="h-10 w-10"
        />
        <div className="w-px flex-1 bg-muted mt-2" />
      </div>

      <form
        onSubmit={handleSubmit(handleCreatePost)}
        className="flex-1 rounded-xl border p-4 shadow-sm bg-card space-y-3"
      >
        <h2 className="text-base font-semibold">Compartilhe algo</h2>
        <div>
          <Input
            placeholder="Título do post"
            disabled={isPending}
            {...register("title")}
          />
          <FormValidationError message={errors.title?.message} />
        </div>
        <div>
          <Textarea
            placeholder="No que você está pensando?"
            className="min-h-[120px] resize-none"
            disabled={isPending}
            {...register("content")}
          />
          <FormValidationError message={errors.content?.message} />
        </div>
        <div className="flex justify-end">
          <Button
            type="submit"
            size="icon"
            variant="ghost"
            disabled={isPending}
          >
            {isPending ? (
              <Loader2 className="animate-spin h-4 w-4" />
            ) : (
              <Send className="h-4 w-4" />
            )}
          </Button>
        </div>
      </form>
    </div>
  );
}
