import { followUser } from "@/api/follow-user";
import { unfollowUser } from "@/api/unfollow-user";
import { Button } from "@/components/ui/button";
import { UserAvatar } from "@/components/user-avatar";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BadgeCheck, Pencil } from "lucide-react";
import { useState } from "react";
import { UpdateProfileModal } from "./update-profile-modal";

interface ProfileHeaderCardProps {
  name: string;
  username: string;
  followers: number;
  following: number;
  isMyProfile: boolean;
  followedByMe?: boolean;
  followingMe?: boolean;
}

export function ProfileHeaderCard({
  name,
  username,
  followers,
  following,
  isMyProfile,
  followedByMe = false,
  followingMe = false,
}: ProfileHeaderCardProps) {
  const [isOpen, setIsOpen] = useState(false);
  const queryClient = useQueryClient();
  const [isFollowing, setIsFollowing] = useState(followedByMe);

  const followMutation = useMutation({
    mutationFn: () => followUser({ username }),
    onSuccess: () => {
      setIsFollowing(true);
      queryClient.invalidateQueries({ queryKey: ["profile", username] });
    },
  });

  const unfollowMutation = useMutation({
    mutationFn: () => unfollowUser({ username }),
    onSuccess: () => {
      setIsFollowing(false);
      queryClient.invalidateQueries({ queryKey: ["profile", username] });
    },
  });

  const handleToggleFollow = () => {
    if (isFollowing) {
      unfollowMutation.mutate();
    } else {
      followMutation.mutate();
    }
  };

  return (
    <>
      <UpdateProfileModal open={isOpen} onOpenChange={setIsOpen} />
      <div className="rounded-xl border shadow-sm bg-card overflow-hidden">
        <div className="h-36 w-full bg-muted relative" />

        <div className="flex flex-col px-6 pt-4 pb-6 -mt-14">
          <div className="flex justify-between items-center relative">
            <UserAvatar
              name={name}
              username={username}
              className="h-24 w-24 border-4 border-background"
            />

            {isMyProfile ? (
              <Button
                variant="outline"
                size="sm"
                className="absolute top-4 right-4 mt-10"
                onClick={() => setIsOpen(true)}
              >
                <Pencil className="w-4 h-4 mr-2" />
                Editar perfil
              </Button>
            ) : (
              <Button
                variant={isFollowing ? "ghost" : "default"}
                size="sm"
                className="absolute top-4 right-4 mt-10"
                onClick={handleToggleFollow}
              >
                {isFollowing ? "Deixar de seguir" : "Seguir"}
              </Button>
            )}
          </div>

          <div className="flex flex-col justify-center mt-1">
            <div className="flex items-center gap-1">
              <h2 className="font-semibold text-xl">{name}</h2>
              <BadgeCheck className="w-4 h-4 text-primary" />
            </div>

            <p className="text-muted-foreground text-sm">@{username}</p>

            <div className="text-xs text-muted-foreground flex gap-4 pt-2">
              <span>
                <strong className="text-foreground">{followers}</strong>{" "}
                seguidores
              </span>
              <span>
                <strong className="text-foreground">{following}</strong> seguindo
              </span>
            </div>

            {!isMyProfile && followingMe && (
              <p className="text-xs text-green-500 mt-2">Este usuário te segue</p>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
