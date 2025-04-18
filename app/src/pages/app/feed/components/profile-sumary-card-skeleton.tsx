import { Skeleton } from "@/components/ui/skeleton";

export function ProfileSummaryCardSkeleton() {
  return (
    <div className="rounded-xl border shadow-sm bg-card overflow-hidden text-center min-h-[192px]">
      <Skeleton className="h-16 w-full" />

      <div className="-mt-8 flex justify-center">
        <Skeleton className="h-12 w-12 rounded-full border-4 border-card" />
      </div>

      <div className="px-4 pt-2 pb-4 space-y-1">
        <div className="flex justify-center items-center gap-1">
          <Skeleton className="h-5 w-28" />
          <Skeleton className="h-4 w-4 rounded-full" />
        </div>

        <Skeleton className="h-4 w-24 mx-auto" />

        <div className="pt-2 text-xs flex justify-center gap-4">
          <Skeleton className="h-4 w-20" />
          <Skeleton className="h-4 w-20" />
        </div>
      </div>
    </div>
  );
}
