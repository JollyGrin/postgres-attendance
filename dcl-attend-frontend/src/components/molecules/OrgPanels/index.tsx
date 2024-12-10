export const OrgPanels = (props: {
  org: {
    img: string;
    name: string;
  };
  locations: {
    name: "hyperfy" | "dcl";
    location: string;
  }[];
}) => {
  return (
    <div className="flex justify-between">
      <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl">
        <img src={props.org.img} height="40px" width="40px" />
        <p>{props.org.name}</p>
      </div>

      <div className="flex gap-4 font-mono">
        {props.locations?.map((loc) => (
          <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl">
            <img src={`${VERSE[loc.name]?.img}`} height="40px" width="40px" />
            <div>
              <p className="text-xs">{VERSE[loc.name]?.name}</p>
              <p className="text-md">{loc.location}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

const VERSE = {
  hyperfy: {
    img: "logos/hyperfy.png",
    name: "hyperfy",
  },
  dcl: {
    img: "logos/dcl.png",
    name: "Decentraland",
  },
};
