import fetch from "axios";

export const fetchPlayerData = async (
  address?: string,
): Promise<PlayerDataResponse> => {
  try {
    const response = await fetch.get<PlayerDataResponse>(
      `https://peer.decentraland.org/lambdas/profile/${address}`,
    );
    return response.data;
  } catch (err) {
    throw err;
  }
};

export type PlayerDataResponse = {
  timestamp: number & { _timestampBrand: never };
  avatars: PlayerDataAvatar[];
};

export type PlayerDataAvatar = {
  hasClaimedName: boolean;
  description: string;
  tutorialStep: number;
  name: string;
  userId: string;
  email: string;
  ethAddress: string;
  version: number;
  avatar: {
    bodyShape: string;
    wearables: string[];
    emotes: {
      slot: number;
      urn: string;
    }[];
    snapshots: {
      body: string;
      face256: string;
    };
    eyes: {
      color: {
        r: number;
        g: number;
        b: number;
        a: number;
      };
    };
    hair: {
      color: {
        r: number;
        g: number;
        b: number;
        a: number;
      };
    };
    skin: {
      color: {
        r: number;
        g: number;
        b: number;
        a: number;
      };
    };
  };
  muted: string[];
};
