import { useQuery } from "@tanstack/react-query";
import { fetchPlayerData } from "./apiClient";

export const useDclPlayer = (address?: string) => {
  return useQuery({
    queryKey: ["dclplayer", address],
    queryFn: () => fetchPlayerData(address),
    enabled: !!address,
    staleTime: Infinity,
  });
};
