package ladok

import (
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/SUNET/goladok3/ladokmocks"
	"github.com/SUNET/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func mockCertificate(t *testing.T, notBefore, notAfter int, folder string) {
	certPEM, _, privateKeyPEM, _ := ladokmocks.MockCertificateAndKey(t, ladoktypes.EnvIntTestAPI, notBefore, notAfter)

	cryptoParts := map[string][]byte{
		"crt": certPEM,
		"key": privateKeyPEM,
	}

	for ext, data := range cryptoParts {
		if err := os.WriteFile(
			filepath.Join(folder, fmt.Sprintf("%s.%s", "testSchoolName", ext)),
			data,
			0777,
		); err != nil {
			t.Fatal(err)
		}
	}
}

func TestImportCertificate(t *testing.T) {
	tempDir := t.TempDir()
	service, _, _, _ := mockService(t, 200, 0, 100, tempDir)
	type have struct {
		notBefore, notAfter int
		fileType            string
	}
	tts := []struct {
		name string
		have have
		want error
	}{
		{
			name: "",
			have: have{},
			want: nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if err := service.Certificate.importCertificate(t.Context()); err != nil {
				assert.EqualError(t, err, tt.want.Error())
			}

			assert.NotNil(t, service.Certificate.Cert)

		})
	}
}

func TestImportCertificateMissingFiles(t *testing.T) {
	tts := []struct {
		name      string
		setupDir  func(t *testing.T) string
		wantErr   bool
	}{
		{
			name: "empty directory - no cert or key",
			setupDir: func(t *testing.T) string {
				return t.TempDir()
			},
			wantErr: true,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)
			cfg := &model.Cfg{}
			cfg.Ladok.Certificate.Folder = dir

			testLog := logger.New("test", true)

			svc := &Service{
				config:     cfg,
				schoolName: "testSchoolName",
			}
			certService := &CertificateService{
				Service: svc,
				logger:  testLog,
			}
			err := certService.importCertificate(t.Context())
			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestCheckValidTime(t *testing.T) {
	type have struct {
		notBefore, notAfter int
	}
	tts := []struct {
		name string
		have have
		want string
	}{
		{
			name: "OK - 100 days cert",
			have: have{
				notBefore: 0,
				notAfter:  100,
			},
			want: CertOK,
		},
		{
			name: "1 day cert",
			have: have{
				notBefore: 0,
				notAfter:  1,
			},
			want: Cert90DaysWarning,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _, _, _ := mockService(t, 200, tt.have.notBefore, tt.have.notAfter, t.TempDir())

			got, _ := service.Certificate.CheckValidTime(t.Context())

			assert.Equal(t, tt.want, got, "status should be equal")
		})
	}
}

func TestIsCertificateInvalid(t *testing.T) {
	type have struct {
		notBefore, notAfter int
	}
	type want struct {
		b bool
		e error
	}
	tts := []struct {
		name string
		have have
		want want
	}{
		{
			name: "Cert into the future, standard case",
			have: have{
				notBefore: 0,
				notAfter:  2,
			},
			want: want{
				b: false,
				e: nil,
			},
		},
		{
			name: "Cert into the future, standard case",
			have: have{
				notBefore: 0,
				notAfter:  100,
			},
			want: want{
				b: false,
				e: nil,
			},
		},
		{
			name: "Cert in the past, old cert",
			have: have{
				notBefore: -2,
				notAfter:  -1,
			},
			want: want{
				b: true,
				e: model.ErrCertificateNotValid,
			},
		},
		{
			name: "NotBefore is After NotAfter, wrong cert",
			have: have{
				notBefore: 2,
				notAfter:  -2,
			},
			want: want{
				b: true,
				e: model.ErrCertificateNotValid,
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _, _, err := mockService(t, 200, tt.have.notBefore, tt.have.notAfter, t.TempDir())
			if tt.want.e != nil {
				//Not possible to execute without a valid certificate
				assert.Equal(t, tt.want.e, err)
			} else {
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				got := service.Certificate.isCertificateInvalid(t.Context())

				assert.Equal(t, tt.want.b, got)
				if !assert.Equal(t, tt.want.e, err) {
					t.FailNow()
				}
			}
		})
	}
}
