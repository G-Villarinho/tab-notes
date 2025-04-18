import { Skeleton } from "@/components/ui/skeleton";

export function PostCardSkeleton() {
  return (
    <div className="flex gap-4">
      <div className="flex flex-col items-center pt-1">
        <Skeleton className="h-10 w-10 rounded-full" />
        <div className="w-px flex-1 bg-muted mt-2" />
      </div>

      <div className="flex-1 rounded-xl border p-4 shadow-sm bg-card space-y-3">
        <div className="flex justify-between items-start gap-2">
          <div className="flex flex-col gap-1">
            <Skeleton className="h-4 w-24" />
            <Skeleton className="h-3 w-16" />
          </div>
          <Skeleton className="h-4 w-4 rounded" />
        </div>

        <div>
          <Skeleton className="h-4 w-1/2" />
          <Skeleton className="h-3 w-full mt-2" />
          <Skeleton className="h-3 w-5/6 mt-1" />
        </div>

        <Skeleton className="h-8 w-24 mt-2 rounded-md" />
      </div>
    </div>
  );
}
