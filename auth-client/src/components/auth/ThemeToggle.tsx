import { IconButton, Tooltip } from "@mui/material";
import Brightness4Icon from "@mui/icons-material/Brightness4";
import Brightness7Icon from "@mui/icons-material/Brightness7";
import { useAtom } from "jotai";
import { themeModeAtom } from "../../store/atoms";
import { useSystemTheme } from "../../theme/theme";

export function ThemeToggle() {
  const systemMode = useSystemTheme();
  const [colorMode, setColorMode] = useAtom(themeModeAtom);
  const mode = colorMode === "system" ? systemMode : colorMode;

  return (
    <Tooltip title={mode === "dark" ? "Switch to light" : "Switch to dark"}>
      <IconButton
        onClick={() => setColorMode(mode === "dark" ? "light" : "dark")}
        aria-label="Toggle color theme"
      >
        {mode === "dark" ? <Brightness7Icon /> : <Brightness4Icon />}
      </IconButton>
    </Tooltip>
  );
}
