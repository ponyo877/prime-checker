import "@typespec/http";

using Http;
@service(#{ title: "Prime Test Service" })
namespace PrimeTestService;

model PrimeTest {
  id: int32;
  number: string;
  created_at: utcDateTime;
}

model PrimeTestRequest {
  number: string;
}

model PrimeTestList {
  items: PrimeTest[];
}

model Setting {
  record_number_success: boolean;
  primality_test_success: boolean;
  email_send_success: boolean;
  dlq_save_success: boolean;
}

@error
model Error {
  code: int32;
  message: string;
}

@route("/primality-tests")
@tag("PrimeTests")
interface PrimeTests {
  @get get(@path request_id: int32): PrimeTest | Error;
  @get list(): PrimeTestList | Error;
  @post create(@body body: PrimeTestRequest): PrimeTest | Error;
}

@route("/settings")
@tag("Settings")
interface Settings {
  @get get(): Setting | Error;
  @post create(@body body: Setting): Setting | Error;
}
