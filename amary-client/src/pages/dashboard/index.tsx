import { Box, Alert, Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, TablePagination, IconButton, Tooltip } from '@mui/material';
import { Helmet } from 'react-helmet-async';
import { useURL } from '../../controllers/useURL';
import BarChartIcon from '@mui/icons-material/BarChart';
import { Link } from 'react-router-dom';

export function DashboardPage() {
  const { userURL, page, limit, hasNextPage, error, handleChangePage, handleChangeLimit } = useURL()
  return (
    <Container maxWidth="md">
      <Helmet>
        <title>Dashboard | Amary</title>
        <meta name="description" content="Manage your shortened URLs, view expiration dates, and access click analytics on the Amary dashboard." />
        <meta name="robots" content="noindex, nofollow" />
        <meta name="googlebot" content="noindex, nofollow" />
      </Helmet>
      <Box sx={{ my: { xs: 3, sm: 4 } }}>
        <Typography
          variant="h3"
          component="h1"
          gutterBottom
          sx={{
            color: 'primary.main',
            fontWeight: 700,
            mb: { xs: 3, sm: 4 },
            fontSize: { xs: '1.9rem', sm: '2.5rem' },
          }}
        >
          Dashboard
        </Typography>

        <Typography
          variant="h5"
          sx={{
            mt: { xs: 3, sm: 4 },
            mb: 2,
            color: 'secondary.main',
            fontWeight: 600,
            fontSize: { xs: '1.2rem', sm: '1.5rem' },
          }}
        >
          Your Shortened URLs
        </Typography>

        {error ?
          <Alert
            severity="error"
            sx={{
              mt: 2,
            }}
          >
            {error.message}
          </Alert>
          : userURL.length === 0 ? (
            <Paper
              variant="outlined"
              sx={{
                p: 4,
                textAlign: 'center',
                borderRadius: 4,
              }}
            >
              <Typography
                sx={{
                  color: 'text.secondary',
                  fontSize: '1.05rem',
                }}
              >
                No URLs shortened yet.
              </Typography>
            </Paper>
          ) : (
            <TableContainer
              component={Paper}
              sx={{
                borderRadius: 1,
                overflowX: 'auto',
              }}
            >
              <Table sx={{ minWidth: 680 }}>
                <TableHead>
                  <TableRow>
                    <TableCell>Short URL</TableCell>
                    <TableCell>Original URL</TableCell>
                    <TableCell>Expires At</TableCell>
                    <TableCell>Created At</TableCell>
                    <TableCell align="right">Stats</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {userURL
                    .map((url) => (
                      <TableRow
                        key={url.id}
                        sx={{
                          '&:hover': {
                            background: 'rgba(180, 83, 9, 0.06)',
                          },
                        }}
                      >
                        <TableCell sx={{ whiteSpace: 'nowrap' }}>
                          <a
                            href={url.short_url}
                            target="_blank"
                            rel="noopener"
                            style={{
                              color: '#B45309',
                              textDecoration: 'none',
                              fontWeight: 600,
                            }}
                          >
                            {url.short_url}
                          </a>
                        </TableCell>
                        <TableCell sx={{ maxWidth: 200, overflow: 'hidden', textOverflow: 'ellipsis', color: 'text.secondary' }}>
                          {url.url}
                        </TableCell>
                        <TableCell sx={{ color: 'secondary.main', whiteSpace: 'nowrap' }}>
                          {url.expired_at
                            ? new Date(url.expired_at).toLocaleDateString()
                            : 'Never'}
                        </TableCell>
                        <TableCell sx={{ color: 'text.secondary', whiteSpace: 'nowrap' }}>
                          {new Date(url.created_at).toLocaleDateString()}
                        </TableCell>
                        <TableCell align="right">
                          <Tooltip title="View Stats" arrow>
                            <IconButton
                              component={Link}
                              to={`/dashboard/url/${url.code}`}
                              size="small"
                              sx={{
                                color: 'primary.main',
                                '&:hover': {
                                  backgroundColor: 'rgba(180, 83, 9, 0.08)',
                                },
                              }}
                            >
                              <BarChartIcon fontSize="small" />
                            </IconButton>
                          </Tooltip>
                        </TableCell>
                      </TableRow>
                    ))}
                </TableBody>
              </Table>
              <TablePagination
                rowsPerPageOptions={[5, 10, 25]}
                component="div"
                count={-1}
                rowsPerPage={limit}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeLimit}
                labelDisplayedRows={({ from, to }) => {
                  return hasNextPage ? `${from}–${to} of more than ${to}` : `${from}–${to} of ${to}`
                }}
                slotProps={{
                  actions: {
                    nextButton: { disabled: !hasNextPage },
                  },
                }}
                sx={{
                  borderTop: '1px solid',
                  borderTopColor: 'divider',
                  color: 'text.secondary',
                  '& .MuiTablePagination-selectLabel, & .MuiTablePagination-displayedRows': {
                    color: 'text.secondary',
                  },
                  '& .MuiIconButton-root': {
                    color: 'primary.main',
                    '&:hover': {
                      backgroundColor: 'rgba(180, 83, 9, 0.08)',
                    },
                  },
                }}
              />
            </TableContainer>
          )}
      </Box>
    </Container>
  );
}
