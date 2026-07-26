import { useParams, Link as RouterLink } from "react-router-dom";
import {
  Box,
  Container,
  Typography,
  Card,
  CardContent,
  Button,
  Grid,
  Tabs,
  Tab,
  CircularProgress,
  IconButton,
  Alert,
  Tooltip,
} from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import ContentCopyIcon from "@mui/icons-material/ContentCopy";
import TodayIcon from "@mui/icons-material/Today";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import BarChartIcon from "@mui/icons-material/BarChart";
import PieChartIcon from "@mui/icons-material/PieChart";
import MultilineChartIcon from "@mui/icons-material/MultilineChart";
import { useURLStats } from "../../controllers/useURLStats";

export function URLStatsPage() {
  const { id } = useParams<{ id: string }>();
  const { stats, urlDetails, loading, error, copied, handleCopy, deviceTab, handleChangeTab, deviceChartRef, dailyChartRef, stackedChartRef } = useURLStats(id);

  if (loading) {
    return (
      <Box sx={{ display: "flex", justifyContent: "center", alignItems: "center", minHeight: "80vh" }}>
        <CircularProgress sx={{ color: "#c25e00" }} />
      </Box>
    );
  }

  if (error) {
    return (
      <Container maxWidth="md" sx={{ my: 4 }}>
        <Alert severity="error">{error.message || "Failed to load URL statistics."}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ my: 4, px: { xs: 2, sm: 3 } }}>
      {/* Back navigation */}
      <Box sx={{ mb: 3 }}>
        <Button
          component={RouterLink}
          to="/dashboard"
          startIcon={<ArrowBackIcon />}
          sx={{
            color: "#c25e00",
            background: "none",
            boxShadow: "none",
            fontSize: "0.95rem",
            px: 1,
            "&:hover": {
              background: "rgba(194, 94, 0, 0.06)",
              boxShadow: "none",
              transform: "none",
            },
          }}
        >
          Back to Dashboard
        </Button>
      </Box>

      {/* URL Meta details card */}
      {urlDetails && (
        <Card sx={{ mb: 4, position: "relative", overflow: "visible" }}>
          <CardContent sx={{ p: { xs: 2.5, sm: 4 } }}>
            <Grid container spacing={3} sx={{
              alignItems: "center"
            }}>
              <Grid item xs={12} md={9}>
                <Typography
                  variant="subtitle2"
                  sx={{ color: "#8b4513", fontWeight: 700, textTransform: "uppercase", letterSpacing: 1, mb: 0.5 }}
                >
                  Shortened Link Analytics
                </Typography>
                <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 1, flexWrap: "wrap" }}>
                  <Typography
                    variant="h4"
                    component="a"
                    href={urlDetails.short_url}
                    target="_blank"
                    rel="noopener noreferrer"
                    sx={{
                      color: "#c25e00",
                      fontWeight: 700,
                      textDecoration: "none",
                      wordBreak: "break-all",
                      fontFamily: "'Merriweather', 'Georgia', serif",
                      "&:hover": {
                        textDecoration: "underline",
                      },
                    }}
                  >
                    {urlDetails.short_url.replace(/^https?:\/\//, "")}
                  </Typography>
                  <Tooltip title={copied ? "Copied!" : "Copy Link"} arrow>
                    <IconButton
                      onClick={handleCopy}
                      size="small"
                      sx={{
                        color: "#c25e00",
                        border: "1px solid #e8dcc8",
                        bgcolor: "#faf6f0",
                        "&:hover": { bgcolor: "#f5e6d3" },
                      }}
                    >
                      <ContentCopyIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                </Box>
                <Typography
                  variant="body2"
                  sx={{
                    color: "text.secondary",
                    wordBreak: "break-all",
                    display: "-webkit-box",
                    WebkitLineClamp: 2,
                    WebkitBoxOrient: "vertical",
                    overflow: "hidden",
                    mb: 1.5,
                  }}
                >
                  Destination:{" "}
                  <a
                    href={urlDetails.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    style={{ color: "#8b4513", textDecoration: "none" }}
                  >
                    {urlDetails.url}
                  </a>
                </Typography>
                <Typography variant="caption" sx={{ color: "text.secondary" }}>
                  Created on: {new Date(urlDetails.created_at).toLocaleDateString(undefined, {
                    year: "numeric",
                    month: "long",
                    day: "numeric",
                  })}
                </Typography>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      )}

      {/* Summary counters grid */}
      <Grid container spacing={3} sx={{ mb: 4, justifyContent: "center" }}>
        <Grid item xs={12} sm={6}>
          <Card
            sx={{
              background: "linear-gradient(135deg, #ffffff 0%, #fdf9f3 100%)",
              borderLeft: "5px solid #c25e00",
            }}
          >
            <CardContent sx={{ display: "flex", alignItems: "center", p: 3 }}>
              <Box
                sx={{
                  bgcolor: "rgba(194, 94, 0, 0.1)",
                  color: "#c25e00",
                  p: 1.5,
                  borderRadius: 2,
                  mr: 2.5,
                  display: "flex",
                }}
              >
                <TodayIcon fontSize="large" />
              </Box>
              <Box>
                <Typography variant="body2" sx={{ color: "#5d4037", fontWeight: 600, textTransform: "uppercase" }}>
                  Visits Today
                </Typography>
                <Typography variant="h3" sx={{ color: "#3e2723", fontWeight: 700, mt: 0.5 }}>
                  {stats?.today_visit_count}
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card
            sx={{
              background: "linear-gradient(135deg, #ffffff 0%, #fdf9f3 100%)",
              borderLeft: "5px solid #8b4513",
            }}
          >
            <CardContent sx={{ display: "flex", alignItems: "center", p: 3 }}>
              <Box
                sx={{
                  bgcolor: "rgba(139, 69, 19, 0.1)",
                  color: "#8b4513",
                  p: 1.5,
                  borderRadius: 2,
                  mr: 2.5,
                  display: "flex",
                }}
              >
                <CalendarMonthIcon fontSize="large" />
              </Box>
              <Box>
                <Typography variant="body2" sx={{ color: "#5d4037", fontWeight: 600, textTransform: "uppercase" }}>
                  Visits This Week
                </Typography>
                <Typography variant="h3" sx={{ color: "#3e2723", fontWeight: 700, mt: 0.5 }}>
                  {stats?.this_week_count}
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Main Stats Charts Grid */}
      <Grid container spacing={3} sx={{ mb: 4, justifyContent: "center" }}>
        {/* Device breakdown doughnut */}
        <Grid item xs={12} md={5}>
          <Card sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
            <CardContent sx={{ p: 3, flexGrow: 1, display: "flex", flexDirection: "column" }}>
              <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center", mb: 2, flexWrap: "wrap", gap: 1 }}>
                <Typography variant="h6" sx={{ color: "#8b4513", display: "flex", alignItems: "center", gap: 1 }}>
                  <PieChartIcon sx={{ color: "#c25e00" }} /> Device Breakdown
                </Typography>
                <Tabs
                  value={deviceTab}
                  onChange={handleChangeTab}
                  size="small"
                  sx={{
                    minHeight: 0,
                    "& .MuiTabs-indicator": { bgcolor: "#c25e00" },
                    "& .MuiTab-root": {
                      py: 0.5,
                      px: 1.5,
                      minHeight: 0,
                      minWidth: 0,
                      fontSize: "0.8rem",
                      fontWeight: 600,
                      color: "#5d4037",
                      "&.Mui-selected": { color: "#c25e00" },
                    },
                  }}
                >
                  <Tab label="Today" />
                  <Tab label="This Week" />
                </Tabs>
              </Box>
              <Box sx={{ flexGrow: 1, minHeight: 250, position: "relative" }}>
                <canvas ref={deviceChartRef} />
              </Box>
            </CardContent>
          </Card>
        </Grid>

        {/* Daily activity bar chart */}
        <Grid item xs={12} md={7}>
          <Card sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
            <CardContent sx={{ p: 3, flexGrow: 1, display: "flex", flexDirection: "column" }}>
              <Typography variant="h6" sx={{ color: "#8b4513", mb: 3, display: "flex", alignItems: "center", gap: 1 }}>
                <BarChartIcon sx={{ color: "#c25e00" }} /> Daily Visits (This Week)
              </Typography>
              <Box sx={{ flexGrow: 1, minHeight: 250, position: "relative" }}>
                <canvas ref={dailyChartRef} />
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Stacked multi-series chart */}
      <Card sx={{ mb: 4 }}>
        <CardContent sx={{ p: 3 }}>
          <Typography variant="h6" sx={{ color: "#8b4513", mb: 3, display: "flex", alignItems: "center", gap: 1 }}>
            <MultilineChartIcon sx={{ color: "#c25e00" }} /> Device Activity Detail by Day
          </Typography>
          <Box sx={{ minHeight: 300, position: "relative" }}>
            <canvas ref={stackedChartRef} />
          </Box>
        </CardContent>
      </Card>
    </Container>
  );
}
