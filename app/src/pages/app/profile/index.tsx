// src/pages/app/profile-page.tsx
import { Helmet } from "react-helmet-async";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { isAxiosError } from "axios";

import { useAuth } from "@/hooks/use-auth";
import { getProfileByUsername } from "@/api/get-profile-by-username";
import { getPostsByUsername } from "@/api/get-posts-by-username";

import { MainNav } from "@/components/main-nav";
import { SuggestedUsersCard } from "@/components/suggested-users-card";
import { TrendingTopicsCard } from "@/components/trending-topics-card";
import { PostCard } from "@/components/post-card";
import { PostCardSkeleton } from "@/components/post-card-skeleton";

import { ProfileHeaderCard } from "./components/profile-header-card";
import { ProfileHeaderCardSkeleton } from "./components/profile-header-card-skeleton";

export function ProfilePage() {
  const { username } = useParams();
  const { user } = useAuth();

  const isMyProfile = !username || username === user?.username;

  const {
    data: publicProfile,
    isLoading: isLoadingProfile,
    isError: isProfileError,
    error: profileError,
  } = useQuery({
    queryKey: ["profile", username],
    queryFn: () => getProfileByUsername({ username: username! }),
    enabled: !!username && !isMyProfile,
    retry: false,
  });

  const profileNotFound =
    isProfileError &&
    isAxiosError(profileError) &&
    profileError.response?.status === 404;

  const activeProfile = isMyProfile ? user : publicProfile;

  const { data: posts, isLoading: isLoadingPosts } = useQuery({
    queryKey: ["posts", username],
    queryFn: () => getPostsByUsername({ username: username! }),
    enabled: !!username && !!activeProfile,
    retry: false,
  });

  return (
    <>
      <Helmet title={`@${activeProfile?.username ?? "Perfil"}`} />
      <div className="mx-auto w-full max-w-screen">
        <div className="grid grid-cols-1 md:grid-cols-[minmax(0,850px)_1fr] gap-6">
          <section className="space-y-6">
            {/* Header */}
            {profileNotFound ? (
              <div className="border text-destructive bg-destructive/10 rounded-lg px-4 py-6 text-center text-sm font-medium">
                Conta não encontrada 😕
              </div>
            ) : isLoadingProfile ? (
              <ProfileHeaderCardSkeleton />
            ) : (
              activeProfile && (
                <ProfileHeaderCard
                  name={activeProfile.name}
                  username={activeProfile.username}
                  followers={activeProfile.followers}
                  following={activeProfile.following}
                  isMyProfile={isMyProfile}
                  followedByMe={
                    !isMyProfile ? publicProfile?.followedByMe : undefined
                  }
                  followingMe={
                    !isMyProfile ? publicProfile?.followingMe : undefined
                  }
                />
              )
            )}

            {/* Posts */}
            {!activeProfile && !profileNotFound
              ? Array.from({ length: 3 }).map((_, i) => (
                  <PostCardSkeleton key={i} />
                ))
              : isLoadingPosts
              ? Array.from({ length: 3 }).map((_, i) => (
                  <PostCardSkeleton key={i} />
                ))
              : posts?.map((post) => (
                  <PostCard
                    key={post.id}
                    post={{
                      postId: post.id,
                      title: post.title,
                      content: post.content,
                      likes: post.likes,
                      createdAt: post.createdAt,
                      likedByUser: post.likedByUser,
                      authorName: activeProfile!.name,
                      authorUsername: activeProfile!.username,
                    }}
                  />
                ))}
          </section>

          <aside className="hidden md:block">
            <div className="sticky top-10 space-y-4">
              <MainNav />
              <SuggestedUsersCard />
              <TrendingTopicsCard />
            </div>
          </aside>
        </div>
      </div>
    </>
  );
}
