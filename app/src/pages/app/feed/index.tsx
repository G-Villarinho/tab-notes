import { Helmet } from "react-helmet-async";
import { useQuery } from "@tanstack/react-query";

import { PostCard } from "@/components/post-card";
import { ProfileSummaryCard } from "./components/profile-summary-card";
import { MainNav } from "@/components/main-nav";
import { SuggestedUsersCard } from "@/components/suggested-users-card";
import { TrendingTopicsCard } from "@/components/trending-topics-card";
import { PostCardSkeleton } from "@/components/post-card-skeleton";

import { useAuth } from "@/hooks/use-auth";
import { getFeed } from "@/api/get-feed";
import { CreatePostCard } from "./components/create-post-card";

export function FeedPage() {
  const { user } = useAuth();

  if (!user) {
    throw new Error("User not found");
  }

  const { data: posts, isLoading: isLoadingPosts } = useQuery({
    queryKey: ["feed"],
    queryFn: () => getFeed({ offset: 0, limit: 10 }),
    retry: false,
    staleTime: 1000 * 60 * 5, // 5 minutes
  });

  return (
    <>
      <Helmet title="feed" />
      <div className="mx-auto w-full max-w-screen">
        <div className="grid grid-cols-1 md:grid-cols-[1fr_minmax(0,600px)_1fr] gap-6">
          <aside className="hidden md:block">
            <div className="sticky top-10 space-y-4">
              {user && <ProfileSummaryCard user={user} />}
              <MainNav />
            </div>
          </aside>

          <section className="space-y-6">
            {user && <CreatePostCard user={user} />}
            {isLoadingPosts && (
              <div className="space-y-6">
                {Array.from({ length: 3 }).map((_, i) => (
                  <PostCardSkeleton key={i} />
                ))}
              </div>
            )}
            {posts?.map((post) => (
              <PostCard key={post.postId} post={post} />
            ))}
          </section>

          <aside className="hidden lg:block">
            <div className="sticky top-10 space-y-4">
              <SuggestedUsersCard />
              <TrendingTopicsCard />
            </div>
          </aside>
        </div>
      </div>
    </>
  );
}
