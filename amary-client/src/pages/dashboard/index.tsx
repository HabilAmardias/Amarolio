import { Box, Alert, Container, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, TablePagination, IconButton, Tooltip } from '@mui/material';
import { useURL } from '../../controllers/useURL';
import BarChartIcon from '@mui/icons-material/BarChart';
import { Link } from 'react-router-dom';

export function DashboardPage() {
  const { userURL, page, limit, hasNextPage, error, handleChangePage, handleChangeLimit } = useURL()
  return (
    <Container maxWidth="md">
      <Box sx={{ my: 4 }}>
        <Typography
          variant="h3"
          component="h1"
          gutterBottom
          sx={{
            color: '#c25e00',
            fontWeight: 700,
            mb: 4,
          }}
        >
          Dashboard
        </Typography>

        <Typography
          variant="h5"
          sx={{
            mt: 4,
            mb: 2,
            color: '#8b4513',
            fontWeight: 600,
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
            <Typography
              sx={{
                color: 'text.secondary',
                fontSize: '1.1rem',
              }}
            >
              No URLs shortened yet.
            </Typography>
          ) : (
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Short URL</TableCell>
                    <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Original URL</TableCell>
                    <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Expires At</TableCell>
                    <TableCell sx={{ color: '#c25e00', fontWeight: 700 }}>Created At</TableCell>
                    <TableCell sx={{ color: '#c25e00', fontWeight: 700 }} align="right">Stats</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {userURL
                    .map((url) => (
                      <TableRow
                        key={url.id}
                        sx={{
                          '&:hover': {
                            background: '#faf6f0',
                          },
                        }}
                      >
                        <TableCell>
                          <a
                            href={url.short_url}
                            target="_blank"
                            rel="noopener"
                            style={{
                              color: '#c25e00',
                              textDecoration: 'none',
                              // '&:hover': {
                              //   textDecoration: 'underline',
                              // },
                            }}
                          >
                            {url.short_url}
                          </a>
                        </TableCell>
                        <TableCell sx={{ maxWidth: 200, overflow: 'hidden', textOverflow: 'ellipsis', color: 'text.secondary' }}>
                          {url.url}
                        </TableCell>
                        <TableCell sx={{ color: '#8b4513' }}>
                          {url.expired_at
                            ? new Date(url.expired_at).toLocaleDateString()
                            : 'Never'}
                        </TableCell>
                        <TableCell sx={{ color: 'text.secondary' }}>
                          {new Date(url.created_at).toLocaleDateString()}
                        </TableCell>
                        <TableCell align="right">
                          <Tooltip title="View Stats" arrow>
                            <IconButton
                              component={Link}
                              to={`/dashboard/url/${url.code}`}
                              size="small"
                              sx={{
                                color: '#c25e00',
                                '&:hover': {
                                  backgroundColor: 'rgba(194, 94, 0, 0.08)',
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
                  borderTop: '1px solid #e8dcc8',
                  color: '#5d4037',
                  '& .MuiTablePagination-selectLabel, & .MuiTablePagination-displayedRows': {
                    color: '#5d4037',
                  },
                  '& .MuiIconButton-root': {
                    color: '#c25e00',
                    '&:hover': {
                      backgroundColor: 'rgba(194, 94, 0, 0.08)',
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
