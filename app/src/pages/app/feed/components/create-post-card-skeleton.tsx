import { Skeleton } from "@/components/ui/skeleton";

export function CreatePostCardSkeleton() {
  return (
    <div className="flex gap-4">
      <div className="flex flex-col items-center pt-1">
        <Skeleton className="h-10 w-10 rounded-full" />
        <div className="w-px flex-1 bg-muted mt-2" />
      </div>

      <div className="flex-1 rounded-xl border p-4 shadow-sm bg-card space-y-3">
        <Skeleton className="h-5 w-1/4" />
        <Skeleton className="h-10 w-full rounded-md" />
        <Skeleton className="h-24 w-full rounded-md" />
        <div className="flex justify-end">
          <Skeleton className="h-10 w-10 rounded-md" />
        </div>
      </div>
    </div>
  );
}
